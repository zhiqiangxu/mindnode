#include "window.h"
#include <stdbool.h>

int InitSDL() {
    if (SDL_Init(SDL_INIT_VIDEO) < 0) {
        fprintf(stderr, "could not initialize sdl2: %s\n", SDL_GetError());
        return 1;
    }

    return 0;
}

int FillSurfaceWithRGB(SDL_Window* window, Uint8 r, Uint8 g, Uint8 b)
{
    SDL_Surface* screenSurface = SDL_GetWindowSurface(window);
    return SDL_FillRect(screenSurface, NULL, SDL_MapRGB(screenSurface->format, r, g, b));
}

void PumpEvent()
{
    SDL_Event event;
    while (true) {
        if (SDL_PollEvent(&event)) {
            
        }
    }
}