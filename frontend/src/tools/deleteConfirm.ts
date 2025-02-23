import { h } from "vue";
import { ExclamationCircleOutlined } from "@ant-design/icons-vue";
import { useModalsStore } from "@/store/modals/modalsStore.ts";

interface DeleteConfirmOptions {
  title?: string;
  content?: string;
  okText?: string;
  cancelText?: string;
  okType?: "primary" | "default" | "dashed" | "danger" | "link";
  centered?: boolean;
  maskClosable?: boolean;
  onOk?: () => void;
  onCancel?: () => void;
}

export const showConfirm = (options: DeleteConfirmOptions) => {
  const {
    title = "Подтвердите действие",
    content = "",
    okText = "Подтвердить",
    cancelText = "Отмена",
    okType = "danger",
    centered = true,
    maskClosable = true,
    onOk = () => console.log("OK"),
    onCancel = () => console.log("Cancel"),
  } = options;

  const modalsStore = useModalsStore();

  modalsStore.modal.confirm({
    title,
    icon: h(ExclamationCircleOutlined),
    content,
    okText,
    cancelText,
    okType,
    centered,
    maskClosable,
    onOk,
    onCancel,
  });
};
