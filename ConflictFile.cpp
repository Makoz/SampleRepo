#include <iostream>

using namespace std;

void foo(){
	
	foo1();
}

void foo1(){
	cout << "CHANGED" << endl;
}

int main(){
	
	foo();
	
	
	return 0;
}
