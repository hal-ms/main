import axios from 'axios'

export default job;
var job = new Job()
class Job {
    constructor() {
        this.axios = axios.create({
            baseURL: 'https://some-domain.com/api/',
            headers: { 'X-Custom-Header': 'foobar' }
        });
    }
    GetJobs(hoge) {
        return {
            "jobs": [
                {
                    "id": "5be0e55734f22500017252c0",
                    "name": "judge",
                    "display_name": "裁判官",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-06T09:50:31.133+09:00"
                },
                {
                    "id": "5be0e55134f22500017252be",
                    "name": "pianist",
                    "display_name": "ピアニスト",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-06T09:50:25.245+09:00"
                },
                {
                    "id": "5bdf085a34f22500017252bc",
                    "name": "judge",
                    "display_name": "裁判官",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T23:55:22.892+09:00"
                },
                {
                    "id": "5bdf065a34f22500017252ba",
                    "name": "carpenter",
                    "display_name": "大工",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T23:46:50.874+09:00"
                },
                {
                    "id": "5bdf065634f22500017252b8",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T23:46:46.057+09:00"
                },
                {
                    "id": "5bdf065234f22500017252b6",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T23:46:42.35+09:00"
                },
                {
                    "id": "5bdf060834f22500017252b4",
                    "name": "priest",
                    "display_name": "僧侶",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T23:45:28.289+09:00"
                },
                {
                    "id": "5bdf058234f22500017252b2",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T23:43:14.164+09:00"
                },
                {
                    "id": "5bdf057d34f22500017252b0",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T23:43:09.294+09:00"
                },
                {
                    "id": "5bdea46334f22500017252ae",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:48:51.899+09:00"
                },
                {
                    "id": "5bdea45034f22500017252ac",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:48:32.181+09:00"
                },
                {
                    "id": "5bdea32834f22500017252a9",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:43:36.079+09:00"
                },
                {
                    "id": "5bdea28c34f22500017252a7",
                    "name": "priest",
                    "display_name": "僧侶",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:41:00.45+09:00"
                },
                {
                    "id": "5bdea28834f22500017252a5",
                    "name": "pianist",
                    "display_name": "ピアニスト",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:40:56.734+09:00"
                },
                {
                    "id": "5bdea28634f22500017252a3",
                    "name": "judge",
                    "display_name": "裁判官",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:40:54.613+09:00"
                },
                {
                    "id": "5bdea28234f22500017252a1",
                    "name": "carpenter",
                    "display_name": "大工",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:40:50.617+09:00"
                },
                {
                    "id": "5bdea28034f225000172529f",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:40:48.05+09:00"
                },
                {
                    "id": "5bdea27a34f225000172529d",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:40:42.85+09:00"
                },
                {
                    "id": "5bdea06534f225000172529b",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T16:31:49.977+09:00"
                },
                {
                    "id": "5bde892734f2250001725299",
                    "name": "judge",
                    "display_name": "裁判官",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T14:52:39.692+09:00"
                },
                {
                    "id": "5bde88ad34f2250001725297",
                    "name": "carpenter",
                    "display_name": "大工",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T14:50:37.14+09:00"
                },
                {
                    "id": "5bde882f34f2250001725295",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T14:48:31.69+09:00"
                },
                {
                    "id": "5bde757534f2250001725293",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T13:28:37.27+09:00"
                },
                {
                    "id": "5bde755434f2250001725291",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T13:28:04.791+09:00"
                },
                {
                    "id": "5bde74f934f225000172528f",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T13:26:33.983+09:00"
                },
                {
                    "id": "5bde67a334f225000172528d",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T12:29:39.418+09:00"
                },
                {
                    "id": "5bde67a034f225000172528b",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T12:29:36.75+09:00"
                },
                {
                    "id": "5bde644c34f2250001725289",
                    "name": "pianist",
                    "display_name": "ピアニスト",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T12:15:24.827+09:00"
                },
                {
                    "id": "5bde644b34f2250001725287",
                    "name": "priest",
                    "display_name": "僧侶",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T12:15:23.35+09:00"
                },
                {
                    "id": "5bde606234f2250001725285",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T11:58:42.124+09:00"
                },
                {
                    "id": "5bde5e5434f2250001725283",
                    "name": "carpenter",
                    "display_name": "大工",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T11:49:56.572+09:00"
                },
                {
                    "id": "5bde5e4d34f2250001725281",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T11:49:49.283+09:00"
                },
                {
                    "id": "5bde5c0234f225000172527e",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T11:40:02.561+09:00"
                },
                {
                    "id": "5bde553634f225000172527c",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T11:11:02.868+09:00"
                },
                {
                    "id": "5bde4bde34f225000172527a",
                    "name": "programmer",
                    "display_name": "プログラマー",
                    "user_name": "mikaboz",
                    "done": false,
                    "notification": false,
                    "created": "2018-11-04T10:31:10.272+09:00"
                },
                {
                    "id": "5bd64e2234f225000172526b",
                    "name": "carpenter",
                    "display_name": "大工",
                    "user_name": "testman",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-29T09:02:42.601+09:00"
                },
                {
                    "id": "5bd64d8e34f2250001725269",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "testman",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-29T09:00:14.598+09:00"
                },
                {
                    "id": "5bd5344934f2250001725265",
                    "name": "carpenter",
                    "display_name": "大工",
                    "user_name": "og3",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-28T13:00:09.896+09:00"
                },
                {
                    "id": "5bcff4fb34f2250001725263",
                    "name": "pianist",
                    "display_name": "ピアニスト",
                    "user_name": "kouki",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-24T13:28:43.515+09:00"
                },
                {
                    "id": "5bcff46734f2250001725261",
                    "name": "pianist",
                    "display_name": "ピアニスト",
                    "user_name": "kouki",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-24T13:26:15.15+09:00"
                },
                {
                    "id": "5bcfe2e9bdfe920001efca1a",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "まっき",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-24T12:11:37.93+09:00"
                },
                {
                    "id": "5bcfd69fbef8b30811977b44",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "まっき",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-24T11:19:11.639+09:00"
                },
                {
                    "id": "5bcfd68ebef8b30811977b42",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "まっき",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-24T11:18:54.288+09:00"
                },
                {
                    "id": "5bcfd646bef8b307c14ef189",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "まっき",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-24T11:17:42.159+09:00"
                },
                {
                    "id": "5bcfd5c7bef8b3079e34f1e2",
                    "name": "cook",
                    "display_name": "料理人",
                    "user_name": "まっき",
                    "done": false,
                    "notification": false,
                    "created": "2018-10-24T11:15:35.758+09:00"
                }
            ]
        }
    }
}