import { getEmitter } from './mitt.js'
import { get_text, get_json } from './fetch.js'
import { local_ws } from './ws.js'

let emitter = getEmitter();

const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY3RpdmUiOiJUIiwidW5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwibmFtZSI6ImFkbWluIiwicGFzc3dvcmQiOiJwYTU1dzByZEBXSVNNRUQiLCJyZWd0aW1lIjoiMjAyMi0wMi0yOFQwNDozMzo0NFoiLCJwaG9uZSI6IiIsImFkZHIiOiIiLCJyb2xlIjoiIiwibGV2ZWwiOiIiLCJleHBpcmUiOiIiLCJuYXRpb25hbGlkIjoiIiwiZ2VuZGVyIjoiIiwicG9zaXRpb24iOiIiLCJ0aXRsZSI6IiIsImVtcGxveWVyIjoiIiwidGFncyI6IiIsImF2YXRhcnR5cGUiOiIiLCJBdmF0YXIiOm51bGwsImV4cCI6MTY0NjI4MjAyNH0.r_6JvJYxWwhOzu59RPfVk57snqLHfnxB-OLkeZE__mg'

export default {

    setup() {
        let myInput = Vue.ref("");
        let imgSrc = Vue.ref("");
        let ws_str = Vue.ref("");
        let resp_str = Vue.ref("");

        // timer example
        let myTimer = setInterval(
            () => {
                let timer_str = Vue.ref("");
                timer_str.value = (new Date()).toLocaleTimeString();
                // send to other app, 'app1' is sender name 
                emitter.emit('app1', timer_str.value + " @ " + myInput.value);
                console.log(myInput.value);
            },
            10000,
        );
        // clearInterval(myTimer);

        /////////////////////////////////////

        // web socket example
        let ws = local_ws("ws/msg"); // hook ws, must be registered in server reg_api
        ws.onopen = function () {
            console.log('ws connected')
        }
        ws.onmessage = function (evt) {
            console.log('ws onmessage', evt.data)
            ws_str.value = evt.data;
        }
        // Send back message, then handle following ws messages in 'onmessage'
        // MUST delay some while !!!
        setTimeout(() => { ws.send('Hello, Server. from com_one.js'); }, 1000);

        /////////////////////////////////////

        // fetch example
        const YesNo = () => {
            // 'async function' return channel             
            (async () => {
                const data = await get_json('https://yesno.wtf/api')
                // console.log(data.answer)
                // console.log(data.image)
                imgSrc.value = data.image
            })();
        }

        const LocalAPI = () => {
            // fetch_get must be here, MUST identical to cert SN

            (async () => {
                const data = await get_text('http://192.168.31.157:1323/api/admin/users', token)
                resp_str.value = data;
            })();

            (async () => {
                const data = await get_json('http://192.168.31.157:1323/api/user/avatar', token)
                imgSrc.value = data.src // "data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="
            })();
        }

        return {
            ws_str,
            myInput,
            imgSrc,
            resp_str,
            YesNo,
            LocalAPI,
        };
    },

    template: `
        <h1>ws message: {{ws_str}}</h1>
        <input v-model="myInput" placeholder="input">
        <br>
        <button class="mybutton" @click="YesNo">YesNoAPI</button>
        <br>
        <img :src="imgSrc" alt="YES/NO IMAGE" width="320" height="240"/>   
        <br>
        <button class="mybutton" @click="LocalAPI">LocalAPI</button>  
        <br>
        <p>response from local API: {{resp_str}}</p>
        <hr>
    `,
};