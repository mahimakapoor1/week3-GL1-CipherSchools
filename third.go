package main
var key int
func GetKey(){
	key=9;
}
type Val struct{
	i int
}
func (v *Val)GetVal()int{
	v.i=10
	return v.i
}
func get(){
	vv:=Val{}
	vv.GetVal()
	ww:=&Val{}
	ww.GetVal()
}