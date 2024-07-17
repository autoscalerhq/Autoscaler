import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '30s', target: 20 }, // Ramp-up to 20 users over 30 seconds
        { duration: '1m', target: 20 },  // Stay at 20 users for 1 minute
        { duration: '30s', target: 0 },  // Ramp-down to 0 users over 30 seconds
    ],
};

export default function () {
    let res = http.get('https://test-api.k6.io/public/crocodiles/');
    check(res, {
        'status is 200': (r) => r.status === 200,
    });
    sleep(1); // Simulate user think time of 1 second
}