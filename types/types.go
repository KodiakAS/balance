package types

type Item struct {
	Data   string `json:"data"`
	Factor int    `json:"factor"`
}

type Items struct {
	Items ItemList `json:"items"`
}

type Bucket struct {
	Items  ItemList
	Factor int
}

type ItemList []Item
type BucketList []Bucket

func (a ItemList) Len() int {
	return len(a)
}

func (a ItemList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ItemList) Less(i, j int) bool {
	return a[i].Factor < a[j].Factor
}

func (a BucketList) Len() int {
	return len(a)
}

func (a BucketList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a BucketList) Less(i, j int) bool {
	return a[i].Factor < a[j].Factor
}

func (b *Bucket) AddItem(i Item) {
	b.Items = append(b.Items, i)
	b.Factor += i.Factor
}
