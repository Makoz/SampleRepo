#include <iostream>

using namespace std;

void foo(){
	cout << "Initial 2" << endl;
	foo1();
}

void foo1(){
	cout << "CHANGED" << endl;
}

int main(){
	
	foo();
	
	
	return 0;
}
