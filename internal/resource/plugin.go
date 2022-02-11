package resource

import (
	"fmt"

	v1 "github.com/kong/koko/internal/gen/grpc/kong/admin/model/v1"
	"github.com/kong/koko/internal/model"
	"github.com/kong/koko/internal/model/json/generator"
	"github.com/kong/koko/internal/model/json/validation"
	"github.com/kong/koko/internal/model/json/validation/typedefs"
	"github.com/kong/koko/internal/plugin"
)

const (
	TypePlugin = model.Type("plugin")
)

var validator plugin.Validator

func SetValidator(v plugin.Validator) {
	validator = v
}

func NewPlugin() Plugin {
	return Plugin{
		Plugin: &v1.Plugin{},
	}
}

type Plugin struct {
	Plugin *v1.Plugin
}

func (r Plugin) ID() string {
	if r.Plugin == nil {
		return ""
	}
	return r.Plugin.Id
}

func (r Plugin) Type() model.Type {
	return TypePlugin
}

func (r Plugin) Resource() model.Resource {
	return r.Plugin
}

func (r Plugin) Validate() error {
	err := validation.Validate(string(TypePlugin), r.Plugin)
	if err != nil {
		return err
	}
	return validator.Validate(r.Plugin)
}

func (r Plugin) ProcessDefaults() error {
	err := validator.ProcessDefaults(r.Plugin)
	return err
}

func (r Plugin) Indexes() []model.Index {
	serviceID, routeID := "", ""
	if r.Plugin.Service != nil {
		serviceID = r.Plugin.Service.Id
	}
	if r.Plugin.Route != nil {
		routeID = r.Plugin.Route.Id
	}
	uniqueValue := fmt.Sprintf("%s.%s.%s", r.Plugin.Name,
		serviceID, routeID)

	res := []model.Index{
		{
			Name: "unique-plugin-per-entity",
			// TODO(hbagdi): needs IndexUniqueMulti for multiple fields?
			Type:  model.IndexUnique,
			Value: uniqueValue,
			// TODO(hbagdi): maybe needs FieldNames?
			FieldName: "",
		},
	}
	if r.Plugin.Route != nil {
		res = append(res, model.Index{
			Name:        "route_id",
			Type:        model.IndexForeign,
			ForeignType: TypeRoute,
			FieldName:   "route.id",
			Value:       r.Plugin.Route.Id,
		})
	}
	if r.Plugin.Service != nil {
		res = append(res, model.Index{
			Name:        "service_id",
			Type:        model.IndexForeign,
			ForeignType: TypeService,
			FieldName:   "service.id",
			Value:       r.Plugin.Service.Id,
		})
	}
	return res
}

func init() {
	err := model.RegisterType(TypePlugin, func() model.Object {
		return NewPlugin()
	})
	if err != nil {
		panic(err)
	}

	const maxProtocols = 8
	pluginSchema := &generator.Schema{
		Type: "object",
		Properties: map[string]*generator.Schema{
			"id":         typedefs.ID,
			"name":       typedefs.Name,
			"created_at": typedefs.UnixEpoch,
			"updated_at": typedefs.UnixEpoch,
			"enabled": {
				Type: "boolean",
			},
			"tags": typedefs.Tags,
			"protocols": {
				Type:     "array",
				Items:    typedefs.Protocol,
				MaxItems: maxProtocols,
			},
			"config": {
				Type:                 "object",
				AdditionalProperties: &truthy,
			},
			"service": typedefs.ReferenceObject,
			"route":   typedefs.ReferenceObject,
		},
		AdditionalProperties: &falsy,
		Required: []string{
			"name",
		},
	}
	err = generator.Register(string(TypePlugin), pluginSchema)
	if err != nil {
		panic(err)
	}
}
