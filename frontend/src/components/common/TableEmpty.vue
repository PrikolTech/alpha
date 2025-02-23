<template>
  <div
    class="container-empty"
    :style="{
      height: fullSize && !tableIsError ? '100%' : '',
    }"
  >
    <empty-block
      :desc="loadingText"
      v-if="tableIsLoading && !hideText"
      :class="{ fullSize: fullSize }"
    />
    <alert-error-load
      v-else-if="tableIsError"
      :text="errorText"
      :tryAgain="tryAgain"
      @clickTryAgainBtn="clickTryAgainHandle"
    />
    <empty-block :desc="emptyText" v-else-if="tableIsEmpty" />
    <empty-block :desc="emptyFilterText" v-else-if="tableIsFilterEmpty" />
    <spinner-absolute
      v-if="spinnerIsVisible && tableIsLoading"
      :background="background"
      :size="size"
      :style="{
        position: spinnerIsFixed ? 'fixed' : 'absolute',
      }"
    />
  </div>
</template>

<script setup lang="ts">
import AlertErrorLoad from "@/components/common/alerts/AlertErrorLoad.vue";
import EmptyBlock from "@/components/common/EmptyBlock.vue";
import SpinnerAbsolute from "@/components/common/SpinnerAbsolute.vue";

interface Props {
  tableIsLoading?: boolean;
  tableIsError?: boolean;
  tableIsEmpty?: boolean;
  tableIsFilterEmpty?: boolean;
  tryAgain?: boolean;
  loadingText?: string;
  errorText?: string;
  emptyText?: string;
  emptyFilterText?: string;
  clickTryAgainHandle?: any;
  spinnerIsVisible?: boolean;
  spinnerIsFixed?: boolean;
  fullSize?: boolean;
  hideText?: boolean;
  background?: boolean;
  size?: "small" | "default" | "large";
}

withDefaults(defineProps<Props>(), {
  tryAgain: false,
  errorText: "Произошла ошибка при загрузке данных",
  loadingText: "Загрузка...",
  emptyText: "Список пуст",
  emptyFilterText: "По вашему запросу ничего не найдено",
});
</script>

<style scoped lang="scss">
.fullSize {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.container-empty {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  position: relative;
}
</style>
