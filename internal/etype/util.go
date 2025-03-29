package etype

func CreateBaseEnumAndSetMap(prefix PrefixType, code int, desc string) BaseEnum {
	baseenum := NewBaseEnum(code, desc)
	Set(prefix, baseenum)
	return *baseenum
}
