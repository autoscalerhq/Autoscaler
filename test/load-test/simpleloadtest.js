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
    let res = http.get('http://localhost:8888/health-check');
    check(res, {
        'status is 200': (r) => r.status === 200,
    });
    if (res.status !== 200){
        console.error("status is: ", res.status)
        console.error('Response Headers:');
        for (const [key, value] of Object.entries(res.headers)) {
            if (key.startsWith('X')) {
                console.error(`${key}: ${value}`);
            }
        }
    }
    sleep(1); // Simulate user think time of 1 second
}