#ifndef TWAIN_WRAPPER_H
#define TWAIN_WRAPPER_H

#ifdef __cplusplus
extern "C" {
#endif

int twain_init();
int twain_select_source();
int twain_enable_adf(int enable);
int twain_set_dpi(int dpi);
int twain_set_color(int color_mode);  // 0=Gray, 1=Color
int twain_acquire(const char* output_dir);  // Returns number of pages
void twain_exit();

#ifdef __cplusplus
}
#endif

#endif