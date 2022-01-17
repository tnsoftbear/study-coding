// Created by Hidden (hidden@undernet.org)

#include <stdio.h>
#include <stdlib.h>
#include <string.h>


char* concat(char *buf, char *s1, char *s2)
{
	// Copy s1 to buf
	strcpy(buf, s1);
	// Append s2 to s1 into buf
	strcat(buf, s2);
	return buf;
}


int main(int argc, char **argv)
{
	char buf[256];
	char buf1[128];
	char buf2[128];

	if (argc != 3)
		return 0;

	// Copy argv[1] to buf1 and argv[2] to buf2
	strncpy(buf1, argv[1], sizeof(buf1));
	strncpy(buf2, argv[2], sizeof(buf2));

	concat(buf, buf2, buf1);
	printf("String result: %s\n", buf);
	return 0;
}
