import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import * as Sentry from "@sentry/browser";
import { BrowserTracing } from "@sentry/tracing";
import { uuid4 } from '@sentry/utils';

export const globalStorage = {
  traceId: ""
}

Sentry.init({
  dsn: "",
  integrations:[
    new BrowserTracing({
      tracePropagationTargets: ['localhost'],
      beforeNavigate: context => {
        globalStorage.traceId = uuid4()
        return {
          ...context,
          traceId: globalStorage.traceId,
          name: window.location.href,
        }
      }
    })
  ],
  tracesSampleRate: 1.0,
});

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
