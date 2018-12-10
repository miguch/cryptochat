import Vue from 'vue';

import BootstrapVue from 'bootstrap-vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

import HelloComponent from './hello.vue';
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
        <hello-component :name="name" :initialEnthusiasm="5" />
    </div>
    `,
    data: { name: cr.GetPublicKey() },
    components: {
        HelloComponent
    }
});



