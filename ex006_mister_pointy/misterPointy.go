package misterPointy

func add(a int, b int) int {
	return a+b
}

func addPointers(a *int, b *int) int {
	return *a + *b
}


func addReturnPointer(a int, b int) *int {
	result := a + b
	return &result
}

func addPointersReturnPointer(a *int, b *int) *int {
	result := *a + *b
	return &result
}

func addModifyArgument(a int, b *int) {
	result := a + *b
	*b = result
}


func addPointerToPointer(a **int, b **int) int {
	result := **a + **b
	return result
}
