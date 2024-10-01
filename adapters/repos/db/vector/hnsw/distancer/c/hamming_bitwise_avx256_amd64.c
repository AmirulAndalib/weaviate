#include <immintrin.h>
#include <nmmintrin.h>
#include <stdint.h>
#include <stdlib.h>
#include <sys/types.h>

static inline uint64_t popcnt_AVX2_lookup(__m256i *vec,
                                          uint64_t *popcnt_constants,
                                          uint8_t *lookup64bit) {

  size_t i = 0;
  // const __m256i lookup = _mm256_setr_epi8(
  //     /* 0 */ 0, /* 1 */ 1, /* 2 */ 1, /* 3 */ 2,
  //     /* 4 */ 1, /* 5 */ 2, /* 6 */ 2, /* 7 */ 3,
  //     /* 8 */ 1, /* 9 */ 2, /* a */ 2, /* b */ 3,
  //     /* c */ 2, /* d */ 3, /* e */ 3, /* f */ 4,

  //     /* 0 */ 0, /* 1 */ 1, /* 2 */ 1, /* 3 */ 2,
  //     /* 4 */ 1, /* 5 */ 2, /* 6 */ 2, /* 7 */ 3,
  //     /* 8 */ 1, /* 9 */ 2, /* a */ 2, /* b */ 3,
  //     /* c */ 2, /* d */ 3, /* e */ 3, /* f */ 4);
  const __m256i lookup = _mm256_loadu_si256((__m256i *)lookup64bit);

  // const __m256i low_mask = _mm256_set1_epi8(0x0f);
  const __m256i low_mask = _mm256_set1_epi64x(popcnt_constants[4]);

  __m256i acc = _mm256_setzero_si256();

  const __m256i lo = _mm256_and_si256(*vec, low_mask);
  const __m256i hi = _mm256_and_si256(_mm256_srli_epi16(*vec, 4), low_mask);
  const __m256i popcnt1 = _mm256_shuffle_epi8(lookup, lo);
  const __m256i popcnt2 = _mm256_shuffle_epi8(lookup, hi);
  __m256i local = _mm256_setzero_si256();
  local = _mm256_add_epi8(local, popcnt1);
  local = _mm256_add_epi8(local, popcnt2);

  acc = _mm256_add_epi64(acc, _mm256_sad_epu8(local, _mm256_setzero_si256()));

  // acc = _mm256_add_epi64(acc, _mm256_sad_epu8(local,
  // _mm256_setzero_si256()));

  uint64_t result = 0;

  result += (uint64_t)(_mm256_extract_epi64(acc, 0));
  result += (uint64_t)(_mm256_extract_epi64(acc, 1));
  result += (uint64_t)(_mm256_extract_epi64(acc, 2));
  result += (uint64_t)(_mm256_extract_epi64(acc, 3));

  return result;
}

static inline uint64_t popcnt_64bit(uint64_t *src, uint64_t *popcnt_constants) {
  uint64_t x = *src;
  x = (x & popcnt_constants[0]) + ((x >> 1) & popcnt_constants[0]);
  x = (x & popcnt_constants[1]) + ((x >> 2) & popcnt_constants[1]);
  x = (x & popcnt_constants[2]) + ((x >> 4) & popcnt_constants[2]);
  return (x * popcnt_constants[3]) >> 56;
}

uint64_t popcnt_lookup_64bit(uint8_t *data, uint64_t *n, uint64_t *lookup64bit,
                             uint64_t *popcnt_constants) {
  uint64_t result = 0;
  uint64_t i = 0;
  uint64_t limit = *n - 7;

  // Unroll the loop 8 times
  while (i < limit) {
    uint64_t sum = popcnt_64bit((uint64_t *)data + i, popcnt_constants);
    result += sum;
    i += 8;
  }

  // Handle remaining elements
  while (i < *n) {
    result += lookup64bit[data[i]];
    i++;
  }

  return result;
}

uint64_t popcount(uint64_t *x, uint64_t *lookup64bit) {
  uint8_t *data = (uint8_t *)x;
  uint64_t result = 0;

  result += lookup64bit[data[0]];
  result += lookup64bit[data[1]];
  result += lookup64bit[data[2]];
  result += lookup64bit[data[3]];
  result += lookup64bit[data[4]];
  result += lookup64bit[data[5]];
  result += lookup64bit[data[6]];
  result += lookup64bit[data[7]];

  return result;
}

void hamming_bitwise_256(uint64_t *a, uint64_t *b, uint64_t *res, long *len,
                         uint8_t *lookup64bit, uint64_t *popcnt_constants) {

  int n = *len;
  uint64_t sum = 0;

  // fast path for small dimensions
  if (n < 8) {
    do {
      uint64_t xor = a[0] ^ b[0];
      sum += popcount(&xor, lookup64bit);

      n--;
      a++;
      b++;
    } while (n);

    *res = sum;
    return;
  }

  __m256i zeros_256 = _mm256_setzero_si256();

  size_t size = 256 / 8;

  while (n >= 16) {
    __m256i a_vec0 = _mm256_loadu_si256((__m256i const *)a);
    __m256i a_vec1 = _mm256_loadu_si256((__m256i const *)(a + 4));
    __m256i a_vec2 = _mm256_loadu_si256((__m256i const *)(a + 8));
    __m256i a_vec3 = _mm256_loadu_si256((__m256i const *)(a + 12));

    __m256i b_vec0 = _mm256_loadu_si256((__m256i const *)b);
    __m256i b_vec1 = _mm256_loadu_si256((__m256i const *)(b + 4));
    __m256i b_vec2 = _mm256_loadu_si256((__m256i const *)(b + 8));
    __m256i b_vec3 = _mm256_loadu_si256((__m256i const *)(b + 12));

    __m256i cmp_result_1 = _mm256_xor_si256(a_vec0, b_vec0);
    __m256i cmp_result_2 = _mm256_xor_si256(a_vec1, b_vec1);
    __m256i cmp_result_3 = _mm256_xor_si256(a_vec2, b_vec2);
    __m256i cmp_result_4 = _mm256_xor_si256(a_vec3, b_vec3);

    uint64_t *p1 = (uint64_t *)&cmp_result_1;
    uint64_t *p2 = (uint64_t *)&cmp_result_2;
    uint64_t *p3 = (uint64_t *)&cmp_result_3;
    uint64_t *p4 = (uint64_t *)&cmp_result_4;

    // sum += popcnt_64bit(p1, popcnt_constants) +
    //        popcnt_64bit(p1 + 1, popcnt_constants) +
    //        popcnt_64bit(p1 + 2, popcnt_constants) +
    //        popcnt_64bit(p1 + 3, popcnt_constants) +
    //        popcnt_64bit(p2, popcnt_constants) +
    //        popcnt_64bit(p2 + 1, popcnt_constants) +
    //        popcnt_64bit(p2 + 2, popcnt_constants) +
    //        popcnt_64bit(p2 + 3, popcnt_constants) +
    //        popcnt_64bit(p3, popcnt_constants) +
    //        popcnt_64bit(p3 + 1, popcnt_constants) +
    //        popcnt_64bit(p3 + 2, popcnt_constants) +
    //        popcnt_64bit(p3 + 3, popcnt_constants) +
    //        popcnt_64bit(p4, popcnt_constants) +
    //        popcnt_64bit(p4 + 1, popcnt_constants) +
    //        popcnt_64bit(p4 + 2, popcnt_constants) +
    //        popcnt_64bit(p4 + 3, popcnt_constants);
    sum += popcnt_AVX2_lookup(&cmp_result_1, popcnt_constants, lookup64bit) +
           popcnt_AVX2_lookup(&cmp_result_2, popcnt_constants, lookup64bit) +
           popcnt_AVX2_lookup(&cmp_result_3, popcnt_constants, lookup64bit) +
           popcnt_AVX2_lookup(&cmp_result_4, popcnt_constants, lookup64bit);

    // sum += popcnt_lookup_64bit(p1, &size, lookup64bit) +
    //        popcnt_lookup_64bit(p2, &size, lookup64bit) +
    //        popcnt_lookup_64bit(p3, &size, lookup64bit) +
    //        popcnt_lookup_64bit(p4, &size, lookup64bit);

    n -= 16;
    a += 16;
    b += 16;
  }

  while (n >= 4) {
    __m256i a_vec0 = _mm256_loadu_si256((__m256i const *)a);
    __m256i b_vec0 = _mm256_loadu_si256((__m256i const *)b);

    __m256i cmp_result_1 = _mm256_xor_si256(a_vec0, b_vec0);

    uint64_t *p1 = (uint64_t *)&cmp_result_1;

    // sum += popcnt_lookup_64bit(p1, &size, lookup64bit, popcnt_constants);
    sum += popcnt_64bit(p1, popcnt_constants) +
           popcnt_64bit(p1 + 1, popcnt_constants) +
           popcnt_64bit(p1 + 2, popcnt_constants) +
           popcnt_64bit(p1 + 3, popcnt_constants);
    n -= 4;
    a += 4;
    b += 4;
  }

  while (n) {
    uint64_t xor = a[0] ^ b[0];
    sum += popcnt_64bit(&xor, popcnt_constants);
    n--;
    a++;
    b++;
  }

  *res = sum;
  return;
}