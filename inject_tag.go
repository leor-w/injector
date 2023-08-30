package injector

import "strings"

const injectTag = "inject"

const (
	tagAliasKey    = "alias"    // 别名
	tagGroupKey    = "scope"    // 分组
	tagOptionalKey = "optional" // 选项
)

// TagOption 标签选项

const (
	TagOptionNotRequired = "NR" // Not Required, 标注该字段在执行注入时, 如果没有找到对应的实例, 则忽略该字段
)

type tagMapper map[string]string

// parseTag 解析 tag
func parseTag(tag string) tagMapper {
	it := make(tagMapper)
	if tag == "" {
		return nil
	}
	if strings.Contains(tag, ",") {
		tags := strings.Split(tag, ",")
		for _, t := range tags {
			if t == "" || !strings.Contains(t, ":") {
				continue
			}
			k, v := parseValue(t)
			it[k] = v
		}
	} else {
		k, v := parseValue(tag)
		it[k] = v
	}
	return it
}

func parseValue(t string) (key, value string) {
	items := strings.Split(t, ":")
	if len(items) == 2 {
		key = items[0]
		value = items[1]
	}
	return
}

func (t tagMapper) getAlias() string {
	return t[tagAliasKey]
}

func (t tagMapper) getGroup() string {
	return t[tagGroupKey]
}

func (t tagMapper) getOptional() string {
	return t[tagOptionalKey]
}

func (t tagMapper) hasAlias() bool {
	return t[tagAliasKey] != ""
}

func (t tagMapper) hasGroup() bool {
	return t[tagGroupKey] != ""
}

func (t tagMapper) hasOptional() bool {
	return t[tagOptionalKey] != ""
}

func (t tagMapper) hasTag() bool {
	return t[tagAliasKey] != "" || t[tagGroupKey] != "" || t[tagOptionalKey] != ""
}
