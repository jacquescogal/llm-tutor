import React from 'react'
import LoginCard from '../components/cards/LoginCard'


const LoginPage = () => {
    const [username, setUsername] = React.useState('')
    const [password, setPassword] = React.useState('')
    const [error, setError] = React.useState('')
  return (
    <div>
        <LoginCard
        username={username}
        password={password}
        error={error}
        setUsername={setUsername}
        setPassword={setPassword}
        setError={setError}
        />
    </div>
  )
}

export default LoginPage