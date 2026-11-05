
#include "wrapper_helper.h"
#include <stdlib.h>
mxSmlReturn_t mxSmlGetLocalAndMultipleRemoteUuidHelper(unsigned char* big_buf, unsigned int* remote_size, unsigned int* uuid_size) {
    unsigned char* local = big_buf;
    unsigned int max_remotes = *remote_size;
    unsigned int initial_size = *uuid_size;
    unsigned char** remotes = (unsigned char**)malloc(max_remotes * sizeof(unsigned char*));
    if (remotes == NULL) {
        return -1;
    }
    for (unsigned int i = 0; i < max_remotes; i++) {
        remotes[i] = big_buf + initial_size + (i * initial_size);
    }
    int ret = mxSmlGetLocalAndMultipleRemoteUuid(local, remotes, remote_size, uuid_size);
    free(remotes);
    return ret;
}