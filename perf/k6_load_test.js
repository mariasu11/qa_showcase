import http from 'k6/http';
import { sleep, check } from 'k6';

export let options = {
  stages: [
    { duration: '30s', target: 20 },
    { duration: '1m',  target: 50 },
    { duration: '30s', target: 0  },
  ],
};

export default function () {
  const url = __ENV.API_ADDR + '/health';
  const res = http.get(url);
  check(res, { 'status was 200': (r) => r.status === 200 });
  sleep(1);
}
