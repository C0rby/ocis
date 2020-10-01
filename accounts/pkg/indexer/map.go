package indexer

// indexMap stores the indexer layout at runtime.

type indexMap map[tName]typeMapping
type tName = string
type fieldName = string

type typeMapping struct {
	pKFieldName    string
	indicesByField map[fieldName][]IndexType
}

func (m indexMap) addIndex(typeName string, pkName string, idx IndexType) {
	if val, ok := m[typeName]; ok {
		val.indicesByField[idx.IndexBy()] = append(val.indicesByField[idx.IndexBy()], idx)
		return
	}
	m[typeName] = typeMapping{
		pKFieldName: pkName,
		indicesByField: map[string][]IndexType{
			idx.IndexBy(): {idx},
		},
	}
}
