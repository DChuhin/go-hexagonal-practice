package mapper

import "github.com/jinzhu/copier"

func Map[Target any, Source any](source *Source) (*Target, error) {
	var target Target
	err := copier.Copy(&target, source)
	if err != nil {
		return nil, err
	}
	return &target, nil
}

func MapSlice[Target any, Source any](source []*Source) ([]*Target, error) {
	target := make([]*Target, len(source))
	err := copier.Copy(&target, &source)
	if err != nil {
		return nil, err
	}
	return target, nil
}
