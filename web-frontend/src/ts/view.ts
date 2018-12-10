import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import HelloComponent from './hello.vue';
import cryptoUtils from "./rsa_crypto";

import blockchain = require("./blockchain");

window.addEventListener('load', function() {
    window.eth = blockchain.getEth();
});

let cr = new cryptoUtils("1234", "1234");
let chatV = new Vue({
    el: "#chat-area",
    template: `
    <div>
        Name: <input v-model="name" type="text">
        <hello-component :name="name" :initialEnthusiasm="5" />
    </div>
    `,
    data: { name: cr.GetPublicKey() },
    components: {
        HelloComponent
    }
});



