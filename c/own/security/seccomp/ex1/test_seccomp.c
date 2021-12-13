# include <stdio.h>
# include <unistd.h> 
# include <linux/seccomp.h> 
# include <sys/prctl.h> 
# include <fcntl.h>  
# include <stdlib.h>

int main () {
   int fd;
   prctl(PR_SET_SECCOMP, SECCOMP_MODE_STRICT);
   fprintf(stderr, "try open\n");
   fd = open("test_file", O_CREAT);
   fprintf(stderr, "fd = %d", fd);
   exit(0);
}
