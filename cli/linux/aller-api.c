#include <stdlib.h>
#include <stdio.h>

int main(){
	char buffer[100];
	snprintf(buffer, sizeof(buffer), "%s ./run.sh", getenv("SHELL"));
	int status = system(buffer);
	if(status != 0) {
		return status;
	}
	return 0;
}
