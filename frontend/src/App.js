import logo from './logo.svg';
import './App.css';
import * as Sentry from "@sentry/browser";
import { uuid4 } from '@sentry/utils';

function App() {
  const sentryTest = () => {
    // const existingTransaction = Sentry.getCurrentHub().getScope()?.getTransaction();
    // const transaction = existingTransaction ??
    //   Sentry.startTransaction({ name: `local test 2`});
    // for (let i = 0; i < 5; i++) {
    //   const span = transaction.startChild({ op: `span ${i}` });
    //   const child = span.startChild({op : `child ${i}-${i}`});
    //   span.finish();
    //   child.finish();
    // }
    // transaction.finish()

    const traceId = uuid4()
    const existingTransaction = Sentry.getCurrentHub().getScope()?.getTransaction();
    const transaction = existingTransaction ?? Sentry.startTransaction({ name: `并列`});
    transaction.traceId = traceId
    const span = transaction.startChild({ op: `span 1` });
    const child = span.startChild({op : `child 1-1`});
    child.finish();
    fetch('http://localhost:8000/sentry', {
      method: 'get',
      headers: {
        "sentry-trace" : transaction.toTraceparent(),
        "trace-id" : traceId,
      },
    }).then((data) => {
      console.log('Success!');
    }).catch((err) => {
      console.log('Something bad happened');
    }).finally( _ => {
      span.finish()
      transaction.finish()
    })
    const t = Sentry.startTransaction({ name: `another`});
    t.traceId = traceId
    t.finish()
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
