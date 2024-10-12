<script lang="ts" setup>
import {
  Search,

} from 'lucide-vue-next'
import { Users } from "~/api/Users";

import { computed, ref } from 'vue'
import { refDebounced } from '@vueuse/core'
import type { Mail } from '../data/mail'
import AccountSwitcher from './AccountSwitcher.vue'

import Nav, { type LinkProp } from './Nav.vue'
import { cn } from '@/lib/utils'
import { Separator } from '@/components/ui/separator'
import { Input } from '@/components/ui/input'
import {
  Tabs,
  TabsContent,
  TabsList,
  TabsTrigger,
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

const isCollapsed = ref(props.defaultCollapsed)

const searchValue = ref('')
const debouncedSearch = refDebounced(searchValue, 250)



// const unreadMailList = computed(() => filteredMailList.value.filter(item => !item.read))

// const selectedMailData = computed(() => props.mails.find(item => item.id === selectedMail.value))





function onCollapse() {
  isCollapsed.value = true
}

function onExpand() {
  isCollapsed.value = false
}

const selectedUser = ref("");

const users = await Users.getContacts();
import { useAuth } from "~/store/auth";

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
  <TooltipProvider :delay-duration="0">
    <ResizablePanelGroup id="resize-panel-group-1" direction="horizontal" class="h-full max-h-[800px] items-stretch">

      <ResizablePanel id="resize-panel-2" :default-size="defaultLayout[1]" :min-size="25">
        <Tabs default-value="all">
          <div class="flex items-center px-4 py-2">
            <h1 class="text-xl font-bold">
              Contact List
            </h1>
            <TabsList class="ml-auto">

              <TabsTrigger value="all" class="text-zinc-600 dark:text-zinc-200">
                All
              </TabsTrigger>
              <TabsTrigger value="unread" class="text-zinc-600 dark:text-zinc-200">
                Unread
              </TabsTrigger>
            </TabsList>
          </div>
          <Separator />
          <div class="bg-background/95 p-4 backdrop-blur supports-[backdrop-filter]:bg-background/60">
            <form>
              <div class="relative">
                <Search class="absolute left-2 top-2.5 size-4 text-muted-foreground" />
                <Input v-model="searchValue" placeholder="Search" class="pl-8" />
              </div>
            </form>
          </div>
          <TabsContent value="all" class="m-0">
            <ContactList  :users="filteredUsers" v-model:selected-user="selectedUser" />
          
          </TabsContent>
          <!-- <TabsContent value="unread" class="m-0">
            <ContactList :users="users" v-model:selected-mail="selectedMail" :items="unreadMailList" />
          </TabsContent> -->
        </Tabs>
      </ResizablePanel>
      <ResizableHandle id="resiz-handle-2" with-handle />
      <ResizablePanel class="bg-slate-100 " id="resize-panel-3" :default-size="defaultLayout[2]">
        <ChatDisplay v-if="selectedUser" :selectedUser="selectedUser" />
        <div v-else class="grid  h-full justify-center  content-center">
          <p class="bg-slate-200 py-2 px-5 rounded-md w-40 text-center">Select a chat.</p>
        </div>
      </ResizablePanel>
    </ResizablePanelGroup>
  </TooltipProvider>
</template>