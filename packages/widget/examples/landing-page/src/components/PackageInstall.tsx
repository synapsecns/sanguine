import { useState } from 'react'

export const PackageInstall = () => {
  const [activeTab, setActiveTab] = useState('npm')

  const tabContainerStyle = {
    display: 'flex',
    marginBottom: '1px',
  }

  const tabStyle = {
    cursor: 'pointer',
    padding: '10px',
    marginRight: '5px',
    borderBottom: '1px solid transparent',
  }

  const activeTabStyle = {
    ...tabStyle,
    borderBottom: '1px solid rgb(217, 70, 239)',
  }

  return (
    <>
      <div style={tabContainerStyle}>
        <div
          style={activeTab === 'npm' ? activeTabStyle : tabStyle}
          onClick={() => setActiveTab('npm')}
        >
          npm
        </div>
        <div
          style={activeTab === 'yarn' ? activeTabStyle : tabStyle}
          onClick={() => setActiveTab('yarn')}
        >
          yarn
        </div>
      </div>
      <pre>
        {activeTab === 'npm'
          ? 'npm i @synapsecns/widget'
          : 'yarn add @synapsecns/widget'}
      </pre>
    </>
  )
}
