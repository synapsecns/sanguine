import Bridge from 'pages/bridge'
import BridgeChain from 'pages/bridge/chain'
import Fee from 'pages/fee'
import Home from 'pages/index'
import Pool from 'pages/pool'
import PoolChain from 'pages/pool/chain'
import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import 'styles/globals.css'

ReactDOM.render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="bridge/:chain" element={<BridgeChain />} />
      <Route path="bridge" element={<Bridge />} />
      <Route path="pool/:chain" element={<PoolChain />} />
      <Route path="pool" element={<Pool />} />
      <Route path="fee" element={<Fee />} />
    </Routes>
  </BrowserRouter>,
  document.getElementById('root'),
)
