export default defineNuxtConfig({
  modules: ['@nuxtjs/tailwindcss', 'shadcn-nuxt','@pinia/nuxt', '@nuxtjs/color-mode',"nuxt-icon","@vee-validate/nuxt"],
  ssr: false,
  runtimeConfig: {
    public: {
      baseURL: process.env.BASE_URL || 'http://localhost:8000',
    },
  },
  buildModules: [
    // Simple usage
    '@nuxtjs/date-fns',

   
  ],
  imports: {
    dirs: ['types/*.ts', 'store/*.ts', 'types/**/*.ts',"api/*.ts"],
  },
  colorMode: {
    classSuffix: ''
  },
  alias: {
    pinia: "/node_modules/@pinia/nuxt/node_modules/pinia/dist/pinia.mjs"
  },
  shadcn: {
    /**
     * Prefix for all the imported component
     */
    prefix: '',
    /**
     * Directory that the component lives in.
     * @default "./components/ui"
     */
    componentDir: './components/ui'
  }
})