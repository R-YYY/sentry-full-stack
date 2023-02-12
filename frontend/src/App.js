import logo from './logo.svg';
import './App.css';
import * as Sentry from "@sentry/browser";
import { globalStorage } from '.';

function App() {
  const sentryTest = () => {
    const existingTransaction = Sentry.getCurrentHub().getScope()?.getTransaction();
    const transaction = existingTransaction ?? Sentry.startTransaction({
      name: 'localhost:3000/',
      op:'http.client'
    });
    transaction.traceId = globalStorage.traceId
    // const span = transaction.startChild({
    //   op: `http.server`,
    //   description: '/api/sentry'
    // });
    fetch('http://localhost:8000/sentry', {
      method: 'get',
      headers: {
        "sentry-trace" : transaction.toTraceparent(),
      },
    }).then((data) => {
      console.log('Success!');
    }).catch((err) => {
      console.log('Something bad happened');
    }).finally( _ => {
      // span.finish()
      transaction.finish()
    })
  }

  const errorTest = () => {
    throw Error("sentry Test")
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
        <button onClick={errorTest}>
          error测试
        </button>
      </header>
    </div>
  );
}

export default App;
