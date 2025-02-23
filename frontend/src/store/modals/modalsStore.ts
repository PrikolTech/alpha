import { defineStore } from "pinia";
import { Modal } from "ant-design-vue";

export const useModalsStore = defineStore("modalsStore", () => {
  const [modal, contextHolder] = Modal.useModal();

  return { modal, contextHolder };
});
