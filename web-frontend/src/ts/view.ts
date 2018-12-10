import Vue from 'vue';

import BootstrapVue from 'bootstrap-vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

import HelloComponent from './components/hello.vue';
import cryptoUtils from "./rsa_crypto";

import blockchain = require("./blockchain");

window.addEventListener('load', function() {
    window.contract = blockchain.getContract();
});

Vue.use(BootstrapVue);


let cr = new cryptoUtils("1234", "1234");
let chatV = new Vue({
    el: "#chat-area",
    template: `
    <div>
        <hello-component v-if="dis" :initialEnthusiasm="5" />
        <p> {{dis}} </p>
    </div>
    `,
    data: { dis: true },
    components: {
        HelloComponent
    }
});

window.chatV = chatV;



