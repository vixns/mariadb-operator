package adapter

type KubernetesAdapter[K any] interface {
	ToKubernetesType() K
}

func ToKubernetesSlice[KA KubernetesAdapter[K], K any](adapters []KA) []K {
	kubernetesSlice := make([]K, len(adapters))
	for i, a := range adapters {
		kubernetesSlice[i] = a.ToKubernetesType()
	}
	return kubernetesSlice
}