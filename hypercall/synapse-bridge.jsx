/* synapse-bridge.jsx
 * A representative rebuild of the Synapse Bridge app (bridge.synapseprotocol.com)
 * as a second backdrop — proves the announcement rides every Synapse property.
 * Exports to window: SynapseBridge
 */

function Caret() {
  return (
    <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" style={{ opacity: 0.6 }}>
      <path d="M6 9l6 6 6-6" />
    </svg>
  );
}

function ChainPill({ icon, name }) {
  return (
    <button style={{
      display: 'inline-flex', alignItems: 'center', gap: 9, cursor: 'pointer',
      background: 'var(--sy-bg)', border: '1px solid var(--sy-panel)', borderRadius: 10,
      padding: '9px 12px', color: '#fff', fontFamily: 'Matter, sans-serif', fontSize: 15,
    }}>
      <img src={`assets/chains/${icon}.svg`} alt="" width="22" height="22" style={{ borderRadius: '50%' }} />
      {name}
      <Caret />
    </button>
  );
}

function TokenPill({ symbol, icon }) {
  return (
    <button style={{
      display: 'inline-flex', alignItems: 'center', gap: 8, cursor: 'pointer',
      background: 'var(--sy-bg)', border: '1px solid var(--sy-panel)', borderRadius: 999,
      padding: '7px 12px 7px 8px', color: '#fff', fontFamily: 'Matter, sans-serif', fontSize: 15, fontWeight: 500,
    }}>
      <img src={`assets/chains/${icon}.svg`} alt="" width="22" height="22" style={{ borderRadius: '50%' }} />
      {symbol}
      <Caret />
    </button>
  );
}

function BridgeRow({ label, chain, token, amount, sub, faded }) {
  return (
    <div style={{ background: 'var(--sy-bg-light)', border: '1px solid var(--sy-panel)', borderRadius: 14, padding: 16 }}>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: 12 }}>
        <span style={{ fontSize: 13, color: 'var(--sy-secondary)' }}>{label}</span>
        <ChainPill icon={chain.icon} name={chain.name} />
      </div>
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', gap: 12 }}>
        <span style={{ fontSize: 30, fontWeight: 500, color: faded ? 'var(--sy-secondary)' : '#fff', letterSpacing: '-0.02em' }}>{amount}</span>
        <TokenPill symbol={token.symbol} icon={token.icon} />
      </div>
      {sub && <div style={{ marginTop: 8, fontSize: 12.5, color: 'var(--sy-secondary)' }}>{sub}</div>}
    </div>
  );
}

function BridgeWidget() {
  return (
    <div style={{ width: 'min(460px, 100%)', background: '#19171c', border: '1px solid var(--sy-panel)', borderRadius: 20, padding: 18, boxShadow: '0 40px 80px -50px rgba(0,0,0,0.9)' }}>
      <div style={{ display: 'flex', gap: 6, marginBottom: 16 }}>
        {['Bridge', 'Swap'].map((tabb, i) => (
          <span key={tabb} style={{
            fontSize: 15, fontWeight: 500, padding: '7px 14px', borderRadius: 9,
            color: i === 0 ? '#fff' : 'var(--sy-secondary)',
            background: i === 0 ? 'rgba(172,143,255,0.14)' : 'transparent',
            border: i === 0 ? '1px solid rgba(172,143,255,0.3)' : '1px solid transparent',
          }}>{tabb}</span>
        ))}
      </div>

      <div style={{ display: 'flex', flexDirection: 'column', gap: 8, position: 'relative' }}>
        <BridgeRow label="From" chain={{ name: 'Ethereum', icon: 'ethereum' }} token={{ symbol: 'USDC', icon: 'base' }} amount="1,000" sub="Balance: 4,920.50" />
        {/* swap arrow */}
        <div style={{ position: 'absolute', left: '50%', top: '50%', transform: 'translate(-50%,-50%)', zIndex: 2 }}>
          <div style={{ width: 36, height: 36, borderRadius: 10, background: '#19171c', border: '1px solid var(--sy-panel)', display: 'grid', placeItems: 'center', color: 'var(--sy-purple)' }}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M7 4v16M7 20l-3-3M7 20l3-3M17 20V4M17 4l-3 3M17 4l3 3" /></svg>
          </div>
        </div>
        <BridgeRow label="To" chain={{ name: 'HyperEVM', icon: 'hyperliquid' }} token={{ symbol: 'USDC', icon: 'base' }} amount="999.41" faded />
      </div>

      <div style={{ display: 'flex', justifyContent: 'space-between', fontSize: 13, color: 'var(--sy-secondary)', padding: '14px 4px 4px' }}>
        <span>Estimated time</span><span style={{ color: '#fff' }}>~2 min · fee $0.59</span>
      </div>

      <button style={{
        width: '100%', marginTop: 12, height: 52, borderRadius: 12, cursor: 'pointer',
        border: '1px solid var(--sy-purple)', color: '#fff', fontFamily: 'Matter, sans-serif', fontSize: 16, fontWeight: 600,
        background: 'linear-gradient(310deg, rgba(255,0,255,0.25), rgba(172,143,255,0.25))',
      }}>Connect Wallet</button>

      <div style={{ textAlign: 'center', marginTop: 14, fontSize: 12.5, color: 'var(--sy-secondary)' }}>Powered by Synapse</div>
    </div>
  );
}

function SynapseBridge({ cfg, annKey }) {
  const ann = ANNOUNCEMENTS[annKey];
  const place = ann ? ann.placement : null;
  const AnnComp = ann ? ann.Comp : null;

  const slotTop = place === 'top' && AnnComp ? <AnnComp cfg={cfg} /> : null;
  // on an app page, both band variants lead full-width under the nav
  const fullBleed = (place === 'belowNav' || place === 'afterHero') && AnnComp;

  return (
    <div className={cfg.motion === 'off' ? '' : 'motion-on'}
      style={{ background: 'radial-gradient(40% 40% at 50% 0%, rgba(172,143,255,0.06), transparent 70%), #111111', color: 'var(--sy-text)', minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
      <SynapseNav slotTop={slotTop} />

      {fullBleed && <AnnComp cfg={cfg} />}

      <main style={{ flex: 1, display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', gap: 24, padding: '48px clamp(16px,4vw,24px) 120px' }}>
        {place === 'hero' && AnnComp && (
          <div style={{ width: 'min(620px, 100%)' }}><AnnComp cfg={cfg} /></div>
        )}
        <BridgeWidget />
      </main>

      {place === 'fixed' && AnnComp && <AnnComp cfg={cfg} />}
    </div>
  );
}

Object.assign(window, { SynapseBridge });
