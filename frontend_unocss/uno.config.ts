import {
    defineConfig,
    presetIcons,
    presetAttributify,
    presetTypography,
    presetUno,
    presetWebFonts
} from 'unocss';

import extractorSvelte from '@unocss/extractor-svelte'
import { FileSystemIconLoader } from '@iconify/utils/lib/loader/node-loaders';

export default defineConfig({
    extractors: [
        extractorSvelte(),
    ],
    presets: [
        presetUno(),
        presetIcons({
            collections: {
                'icons': FileSystemIconLoader(
                    './static/icons',
                )
            }
        }),
        presetAttributify({
            strict: true,
            ignoreAttributes: [
                "click"
            ]
        }),
        presetTypography(),
        presetWebFonts({
            provider: 'google',
            fonts: {
                robot: ['Roboto Mono'],
                sans: ['Inconsolate', 'Nunito Sans'],
                mono: ['JetBrains Mono'],
            },
        }),
    ]
})