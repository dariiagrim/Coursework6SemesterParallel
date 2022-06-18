#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct MatrixABI {
  uintptr_t rows;
  uintptr_t cols;
  double *data;
} MatrixABI;

struct MatrixABI *matrix_new(uintptr_t rows, uintptr_t cols, double *data);

struct MatrixABI *multiply_async(struct MatrixABI *matrix, struct MatrixABI *other);
