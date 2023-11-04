package model

import (
	"reflect"
	"strings"
)

func (m *Model) SearchContacts(filters ...string) (ContactsSet, error) {
	m.app.Info("running SearchContacts method...")
	defer m.app.Info("finished running SearchContacts method!")

	payload := ContactsSet{}

	if (len(filters) == 1 && filters[0] == "") || len(filters) == 0 || filters == nil {
		m.app.Warning("search contacts called without args returning all contacts")
		return m.contacts, nil
	}

	for _, filter := range filters {
		for _, cont := range m.contacts {
			v := reflect.ValueOf(*cont)
			for i := 0; i < v.NumField(); i++ {
				field := v.Field(i)
				if field.Kind() != reflect.String {
					continue
				}
				fieldValue := field.String()
				if !strings.Contains(strings.ToLower(fieldValue), strings.ToLower(filter)) {
					continue
				}
				payload[cont.ID] = cont
			}
		}
	}

	return payload, nil
}
