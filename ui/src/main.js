import { createApp } from "vue";
import App from "./App.vue";
import { createMetaManager } from "vue-meta";
import "./index.css";
import router from "./router";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faSearch } from "@fortawesome/free-solid-svg-icons";

library.add(faSearch);

createApp(App)
  .use(router)
  .use(createMetaManager())
  .component("font-awesome-icon", FontAwesomeIcon)
  .mount("#app");
