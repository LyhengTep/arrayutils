package arrayutils


func GetArrayElement [K any] (arr []K,index int) *K{
	if len(arr)>index{
		return &arr[index];
	}
	return nil
}