package config

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"

	"github.com/kong/koko/internal/json"
)

type Map map[string]interface{}

type Content struct {
	CompressedPayload []byte
	Hash              string
	GranularHashes    map[string]string
}

type MutatorOpts struct {
	ClusterID string
}

type Mutator interface {
	Name() string
	Mutate(context.Context, MutatorOpts, DataPlaneConfig) error
}

type Loader interface {
	Load(ctx context.Context, clusterID string) (Content, error)
}

type DataPlaneConfig Map

type KongConfigurationLoader struct {
	mutators []Mutator
}

func (l *KongConfigurationLoader) Register(mutator Mutator) error {
	for _, m := range l.mutators {
		if m.Name() == mutator.Name() {
			return fmt.Errorf("mutator '%v' already registered", m.Name())
		}
	}
	l.mutators = append(l.mutators, mutator)
	return nil
}

func (l *KongConfigurationLoader) Load(ctx context.Context, clusterID string) (Content, error) {
	var configTable DataPlaneConfig = map[string]interface{}{}
	for _, m := range l.mutators {
		err := m.Mutate(ctx, MutatorOpts{ClusterID: clusterID},
			configTable)
		if err != nil {
			return Content{}, err
		}
	}

	return ReconfigurePayload(configTable)
}

func ReconfigurePayload(c DataPlaneConfig) (Content, error) {
	configHash, hashes := getGranularHashes(c)
	payload := Map{
		"type":         "reconfigure",
		"config_table": c,
		"config_hash":  configHash,
		"hashes":       hashes,
	}

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	defer writer.Close()

	err := json.Marshaller.NewEncoder(writer).Encode(payload)
	if err != nil {
		return Content{}, fmt.Errorf("json marshal: %v", err)
	}
	err = writer.Close()
	if err != nil {
		return Content{}, fmt.Errorf("gzip failure: %v", err)
	}
	return Content{
		CompressedPayload: buf.Bytes(),
		Hash:              configHash,
		GranularHashes:    hashes,
	}, nil
}

func CompressPayload(payload []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(payload)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func UncompressPayload(payload []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	buf, err := io.ReadAll(r)
	if err != nil {
		_ = r.Close()
		return nil, err
	}
	err = r.Close()
	if err != nil {
		return nil, err
	}
	return buf, nil
}
