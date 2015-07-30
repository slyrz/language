#ifndef BINDING_CLD2_H
#define BINDING_CLD2_H

#include <stdbool.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

struct language {
  const char *code;
  const char *name;
  int reliable;
};

void detect_language(char *buffer, int length, bool is_plain_text, struct language *language);

#ifdef __cplusplus
}
#endif
#endif
