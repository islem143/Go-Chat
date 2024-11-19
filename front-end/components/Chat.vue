<script lang="ts" setup>
import {
  Search,

} from 'lucide-vue-next'
import { Users } from "~/api/Users";

import { computed, ref } from 'vue'
import { refDebounced } from '@vueuse/core'
import type { Mail } from '../data/mail'
import { useAuth } from "~/store/auth";
import { Separator } from '@/components/ui/separator'
import { Input } from '@/components/ui/input'
import {
  Tabs,
  TabsContent,

} from '@/components/ui/tabs'
import { TooltipProvider } from '@/components/ui/tooltip'
import { ResizableHandle, ResizablePanel, ResizablePanelGroup } from '@/components/ui/resizable'

interface MailProps {
  accounts: {
    label: string
    email: string
    icon: string
  }[]
  mails: Mail[]
  defaultLayout?: number[]
  defaultCollapsed?: boolean
  navCollapsedSize: number
}

const props = withDefaults(defineProps<MailProps>(), {
  defaultCollapsed: false,
  defaultLayout: () => [265, 100, 655],
})



const searchValue = ref('')
const debouncedSearch = refDebounced(searchValue, 250)




const selectedUser = ref("");

const users = await Users.getContacts();


const auth = useAuth();
console.log(auth);
const filteredUsers = computed(() => {
  let output: Mail[] = []
  const searchValue = debouncedSearch.value?.trim()
  if (!searchValue) {
    output = users
  }

  else {
    output = users.filter((item) => {
      return item.name.includes(debouncedSearch.value)

    })
  }

  return output
})

</script>

<template>
  <div class="grid grid-cols-2 md:grid-cols-5 p-5">

    <Card class="col-span-2 md:col-span-1">
      <h1 class="text-xl font-bold p-2">
        Contact List
      </h1>



      <div class="bg-background/95 p-4 backdrop-blur supports-[backdrop-filter]:bg-background/60">
        <form>
          <div class="relative">
            <Search class="absolute left-2 top-2.5 size-4 text-muted-foreground" />
            <Input v-model="searchValue" placeholder="Search" class="pl-8" />
          </div>
        </form>
      </div>

      <ContactList :users="filteredUsers" v-model:selected-user="selectedUser" />

    </Card>



    <div class="col-span-2 md:col-span-4  max-h-[800px] ">
      <ChatDisplay v-if="selectedUser" :selectedUser="selectedUser" />
      <div v-else class="grid  h-full justify-center  content-center">
        <p class=" py-2 px-5 rounded-md w-40 text-center">Select a chat.</p>
      </div>
    </div>
  </div>

</template>