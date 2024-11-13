<script lang="ts" setup>

import { accounts, mails } from '../components/data/mail'
</script>

<template>
  <div class="md:hidden">
    <image
      src="/examples/mail-dark.png"
      :width="1280"
      :height="727"
      alt="Mail"
      class="hidden dark:block"
    />
    <image
      src="/examples/mail-light.png"
      :width="1280"
      :height="727"
      alt="Mail"
      class="block dark:hidden"
    />
  </div>
  <div class=" flex-col md:flex">
    <Chat
      :accounts="accounts"
   
   
    />
  </div>
</template>