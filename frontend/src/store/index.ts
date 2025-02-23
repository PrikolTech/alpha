import { createPinia, defineStore } from "pinia";
import type { StateTree } from "pinia";

export const useCommonStore = defineStore("commonStore", () => {});

function storeReset({ store }: StateTree) {
  // Плагин для pinia добавляет функцию "$reset"(при использовании pinia с setup ее нет)
  // клонирование объекта с помощью "JSON.parse" что бы не тянуть библиотеку lodash
  const initialState = JSON.parse(JSON.stringify(store.$state));
  store.$reset = () => store.$patch(JSON.parse(JSON.stringify(initialState)));
}

const store = createPinia();
store.use(storeReset);

export default store;
