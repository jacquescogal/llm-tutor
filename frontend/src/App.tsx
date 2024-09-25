import './App.css'
import { useSelector } from 'react-redux'
import { RootState } from './store/store'
import { useDispatch } from 'react-redux'
import { decrement, increment, incrementByAmount } from './store/counterSlice'
import { changeString } from './store/stringExample'
import Card from './components/Card'
function App() {
  // access state with useSelector, typed with RootState
  const count = useSelector((state: RootState) => state)

  // Use AppDispatch type for dispatch
  const dispatch = useDispatch()
  return (
    <>
    <div>
      {count.counter.value}
      <button onClick={() => dispatch(increment())}>+</button>
      <button onClick={() => dispatch(decrement())}>-</button>
      <button onClick={() => dispatch(incrementByAmount(2))}>+2</button>
    </div>
    <div>
      <Card 
      title={"hello"}
      body=
      {count.hmmInteresting.value}
      buttons={[<button onClick={() => dispatch(changeString("what"))}>Change String</button>,<button onClick={() => dispatch(changeString("what"))}>Change String</button>,<button onClick={() => dispatch(changeString("what"))}>Change String</button>]}/>
      <Card 
      title={"hello"}
      body=
      {count.hmmInteresting.value}
      buttons={[<button onClick={() => dispatch(changeString("what"))}>Change String</button>,<button onClick={() => dispatch(changeString("what"))}>Change String</button>,<button onClick={() => dispatch(changeString("what"))}>Change String</button>]}/>
      <Card 
      title={"hello"}
      body=
      {count.hmmInteresting.value}
      buttons={[<button onClick={() => dispatch(changeString("what"))}>Change String</button>,<button onClick={() => dispatch(changeString("what"))}>Change String</button>,<button onClick={() => dispatch(changeString("what"))}>Change String</button>]}/>
    </div>
    </>
  )
}

export default App
