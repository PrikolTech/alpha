/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_PACKAGE_VERSION: string;
  readonly VITE_API_PORT: string;
  readonly API_PORT: string;
  readonly API_HOSTNAME: string;
  readonly API_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}

declare const PACKAGE_VERSION: string;
