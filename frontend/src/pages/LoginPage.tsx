import React from 'react'
import LoginCard from '../components/cards/LoginCard'
import { CrumbHelperSetRootLink } from '../store/helpers/crumbHelper'


const LoginPage = () => {

  CrumbHelperSetRootLink({name: "Login"});
  return (
    <div className='flex flex-col items-center h-full justify-center'>
        <LoginCard/>
    </div>
  )
}

export default LoginPage