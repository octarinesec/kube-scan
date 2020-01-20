import { debug } from 'utils/logging';

export let API_ENDPOINT = process.env.REACT_APP_API_ENDPOINT || '/';

if (API_ENDPOINT && API_ENDPOINT.length && API_ENDPOINT !== '/' && API_ENDPOINT[API_ENDPOINT.length-1] !== '/') {
  API_ENDPOINT += '/';
}

if ((process.env.NODE_ENV === 'development')) {
  debug('API_ENDPOINT', API_ENDPOINT);
}

export let USER_DEFAULT_NS = '';
export let USER_DEFAULT_NAME = '';
export let USER_DEFAULT_PASSWORD = '';

if ((process.env.NODE_ENV === 'development') && process.env.REACT_APP_AUTOFILL_USER_DETAILS) {
  USER_DEFAULT_NS = process.env.REACT_APP_API_USER_DEFAULT_NS;
  USER_DEFAULT_NAME = process.env.REACT_APP_API_USER_DEFAULT_NAME;
  USER_DEFAULT_PASSWORD = process.env.REACT_APP_API_USER_DEFAULT_PASSWORD;
}

export const API_CALL_TIMEOUT = 25 * 1000; // 25 sec;


export const FETCH_INTERVALS = {
  AUTH_TOKEN: {
    REFRESH: 60 * 1000,               // 1 minute,
    RETRY: 3 * 1000,                  // 3 seconds,
  },
  ACTIVE_REPLICAS: {
    REFRESH: 60 * 1000,               // 1 minute
  },
  ANOMALY_ALERTS: {
    SLOW: 3 * 60 * 1000,              // 3 minutes,
    RAPID: 20 * 1000,                 // 20 seconds,
  },
  BLOCKED_MESSAGES: {
    SLOW: 3 * 60 * 1000,              // 3 minutes,
    RAPID: 20 * 1000,                 // 20 seconds,
  },
  INSIGHTS: {
    REFRESH: 3 * 60 * 1000,           // 3 minutes,
  },
  K8S_CHANGES : {
    SLOW: 10 * 60 * 1000,             // 10 minutes
    RAPID: 30 * 1000,                 // 30 seconds
  },
  K8S_ALERTS : {
    SLOW: 20 * 60 * 1000,             // 20 minutes
    RAPID: 30 * 1000,                 // 30 seconds
  },
  K8S_POLICIES : {
    SLOW: 20 * 60 * 1000,             // 20 minutes
    RAPID: 60 * 1000,                 // 60 seconds
  },
  K8S_CUSTOM_RULES : {
    REFRESH: 20 * 60 * 1000,          // 20 minutes
    RETRY: 5 * 1000,                  // 5 seconds
  },
  KAFKA: {
    FIRST: 1500,                       // 1.5 second
    REGULAR: 61 * 1000,               // 61 seconds
  },
  NETWORK_ALERTS: {
    SLOW: 3 * 60 * 1000,              // 3 minutes,
    RAPID: 20 * 1000,                 // 20 seconds,
    FORCE_ALLOW_AFTER: 20 * 1000,     // 20 seconds
  },
  THREAT_ALERTS: {
    SLOW: 6 * 60 * 1000,              // 6 minutes,
    RAPID: 2 * 60  * 1000,            // 2 minutes,
  },
  THREAT_RULES: {
    SLOW: 5 * 60 * 1000,              // 5 minutes,
    RAPID: 20 * 1000,                 // 20 seconds,
  },
  NETWORK_ACTIVITY_SUMMARY: {
    SLOW: 3 * 60 * 1000,              // 2 minutes,
    RAPID: 20 * 1000,                 // 30 seconds,
  },
  REPORT_CARDS: {
    SLOW: 10 * 60 * 1000,             // 10 minutes
    RAPID: 60 * 1000,                 // 1 minute
  },
  WORKLOAD_DATA: {
    REFRESH: 3 * 60 * 1000,           // 2 minutes,
  },
};

export const DEFAULT_AGGREGATION_THRESHOLD = 10;
export const KAFKA_DEFAULT_AGGREGATION_THRESHOLD = 20;
