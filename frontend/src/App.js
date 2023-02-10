import logo from './logo.svg';
import './App.css';
import * as Sentry from "@sentry/browser";
import { BrowserTracing } from "@sentry/tracing";

Sentry.init({
  dsn: "https://a9bab2e92f48490484543f621223ee89@o4504648733032448.ingest.sentry.io/4504648959852544",

  integrations:[
    new BrowserTracing({
      tracePropagationTargets: ['localhost'],
    })
  ],

  // Set tracesSampleRate to 1.0 to capture 100%
  // of transactions for performance monitoring.
  // We recommend adjusting this value in production
  tracesSampleRate: 1.0,
});

function App() {
  const sentryTest = () => {
    const transaction = Sentry.startTransaction({ name: "test-transaction" });
    const span = transaction.startChild({ op: "sentryTest" }); // This function returns a Span
    // functionCallX

    fetch('http://localhost:8000/sentry',{
      method: 'get',
      headers: {
        'sentry-trace': span.toTraceparent(),
        // 'baggage': serializeBaggage(span.getBaggage())
      }
    }).then((data) => {
      console.log('Success!');
    }).catch((err) => {
      console.log('Something bad happened');
    }).finally(_ => {
      span.finish(); // Remember that only finished spans will be sent with the transaction
      transaction.finish(); // Finishing the transaction will send it to Sentry
    })
  }
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <button onClick={sentryTest}>
          sentry测试
        </button>
      </header>
    </div>
  );
}

export default App;
