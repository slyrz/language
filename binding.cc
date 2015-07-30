#include "binding.h"

#include <cstddef>
#include <cstdio>
#include "cld2/public/compact_lang_det.h"

void detect_language(char* buffer, int length, bool is_plain_text, struct language* language) {
  bool is_reliable;
  double normalized_score3[3];
  int flags = 0;
  int percent3[3];
  int text_bytes;
  int valid_prefix_bytes;

  CLD2::CLDHints cldhints = {NULL, NULL, 0, CLD2::UNKNOWN_LANGUAGE};
  CLD2::Language language3[3];
  CLD2::Language result;
  CLD2::ResultChunkVector resultchunkvector;

  if (language == NULL) return;

  result = CLD2::ExtDetectLanguageSummaryCheckUTF8(
      buffer, length, is_plain_text, &cldhints, flags, language3, percent3, normalized_score3,
      &resultchunkvector, &text_bytes, &is_reliable, &valid_prefix_bytes);

  language->code = CLD2::LanguageCode(result);
  language->name = CLD2::LanguageName(result);
  language->reliable = is_reliable;
}
