<template>
  <a-config-provider :locale="localeConfigurations.ru">
    <router-view />
    <component :is="contextHolder" />
  </a-config-provider>
</template>

<script setup lang="ts">
import { useModalsStore } from "@/store/modals/modalsStore.ts";
import ruRU from "ant-design-vue/es/locale/ru_RU";
import enUS from "ant-design-vue/es/locale/en_US";

const localeConfigurations: Record<string, any> = {
  ru: ruRU,
  en: enUS,
};
const modalsStore = useModalsStore();
const { contextHolder } = modalsStore;
</script>

<style scoped lang="scss"></style>
