/* synapse-page.jsx
 * Faithful rebuild of the Synapse landing page (the backdrop), with the
 * Hypercall announcement injected at the slot its placement calls for.
 * Exports to window: SynapsePage
 */

const SY_NAV = ['About', 'Bridge', 'Swap', 'Pools', 'Stake', '$SYN', 'Solana Bridge'];

const SUPPORT_CARDS = [
  { t: 'Extensible', d: 'Build cross-chain apps that send messages and assets across any supported chain.' },
  { t: 'Secure', d: 'An optimistic security model with a global network of independent verifiers.' },
  { t: 'Generalized', d: 'Move data and value in a single transaction — not just tokens.' },
];
const BRIDGE_CARDS = [
  { t: 'Deep Liquidity', d: 'Swap native assets across chains with the deepest cross-chain liquidity.' },
  { t: 'Wide Support', d: 'Access 20+ chains and the assets that matter, from one interface.' },
  { t: 'Developer Friendly', d: 'Integrate the bridge and router with a few lines of code.' },
];
const CHAINS = [
  ['Ethereum', 'ethereum'], ['Base', 'base'], ['Arbitrum', 'arbitrum'], ['Optimism', 'optimism'],
  ['HyperEVM', 'hyperliquid'], ['Berachain', 'berachain'], ['Avalanche', 'avalanche'],
  ['Polygon', 'polygon'], ['BNB Chain', 'bnb'], ['Blast', 'blast'], ['Unichain', 'unichain'],
];

function CardIcon({ kind }) {
  // simple, on-brand geometric glyphs (no slop)
  const g = 'url(#sy-card-grad)';
  return (
    <svg width="40" height="40" viewBox="0 0 40 40" fill="none">
      <defs>
        <linearGradient id="sy-card-grad" x1="0" y1="0" x2="40" y2="40">
          <stop stopColor="var(--sy-pink)" /><stop offset="1" stopColor="var(--sy-purple)" />
        </linearGradient>
      </defs>
      {kind === 0 && <><rect x="5" y="5" width="13" height="13" rx="3" stroke={g} strokeWidth="2" /><rect x="22" y="22" width="13" height="13" rx="3" stroke={g} strokeWidth="2" /><path d="M18 11h10v11" stroke={g} strokeWidth="2" /></>}
      {kind === 1 && <><path d="M20 4l13 6v9c0 8-5.5 13-13 17-7.5-4-13-9-13-17V10l13-6Z" stroke={g} strokeWidth="2" strokeLinejoin="round" /><path d="M14 20l4.5 4.5L27 16" stroke={g} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" /></>}
      {kind === 2 && <><circle cx="10" cy="20" r="5" stroke={g} strokeWidth="2" /><circle cx="30" cy="10" r="5" stroke={g} strokeWidth="2" /><circle cx="30" cy="30" r="5" stroke={g} strokeWidth="2" /><path d="M14.5 17.5 25.5 12M14.5 22.5 25.5 28" stroke={g} strokeWidth="2" /></>}
    </svg>
  );
}

function SynapseNav({ slotTop }) {
  return (
    <header>
      {slotTop}
      <div style={{ maxWidth: 1440, margin: '0 auto', display: 'flex', justifyContent: 'space-between',
        alignItems: 'center', gap: 16, padding: 'clamp(20px,3vw,30px) clamp(18px,4vw,40px)' }}>
        <img src="assets/synapse-logo.svg" alt="Synapse by Hypercall" style={{ height: 34, width: 'auto' }} />
        <nav style={{ display: 'flex', flexWrap: 'wrap', justifyContent: 'center', gap: 2 }} className="sy-desktop-nav">
          {SY_NAV.map((item, i) => (
            <a key={item} href="#" style={{
              padding: '8px 12px', color: '#fff', whiteSpace: 'nowrap',
              opacity: i === 0 ? 1 : 0.34, fontSize: 15, transition: 'opacity .15s',
            }}
              onMouseEnter={(e) => (e.currentTarget.style.opacity = 1)}
              onMouseLeave={(e) => (e.currentTarget.style.opacity = i === 0 ? 1 : 0.34)}>
              {item}
            </a>
          ))}
        </nav>
        <div style={{ display: 'flex', alignItems: 'center', gap: 10 }} className="sy-desktop-nav">
          <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="rgba(255,255,255,0.7)" strokeWidth="1.5">
            <path strokeLinecap="round" strokeLinejoin="round" d="M12 3c4.97 0 9 4.03 9 9s-4.03 9-9 9-9-4.03-9-9 4.03-9 9-9Zm0 0c2.4 2.2 3.75 5.41 3.75 9S14.4 18.8 12 21m0-18C9.6 5.2 8.25 8.41 8.25 12S9.6 18.8 12 21m-8.25-9h16.5" />
          </svg>
          <button style={{ width: 38, height: 38, borderRadius: 8, border: '1px solid var(--sy-panel)', background: 'transparent', color: '#fff', cursor: 'pointer' }}>
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.5" strokeLinecap="round"><path d="M5 12h.01M12 12h.01M19 12h.01" /></svg>
          </button>
        </div>
      </div>
    </header>
  );
}

function SectionLabel({ children }) {
  return <div style={{ fontFamily: 'var(--mono)', fontSize: 12, letterSpacing: '0.14em', textTransform: 'uppercase', color: 'var(--sy-purple)', marginBottom: 14 }}>{children}</div>;
}

function SupportCard({ icon, t, d }) {
  return (
    <div style={{ maxWidth: 360 }}>
      <div style={{ marginBottom: 14 }}><CardIcon kind={icon} /></div>
      <div style={{ fontSize: 26, color: '#fff', marginBottom: 8, letterSpacing: '-0.01em' }}>{t}</div>
      <div style={{ color: 'var(--sy-secondary)', fontSize: 15, lineHeight: 1.55 }}>{d}</div>
    </div>
  );
}

function SynapsePage({ cfg, annKey }) {
  const ann = ANNOUNCEMENTS[annKey];
  const place = ann ? ann.placement : null;
  const AnnComp = ann ? ann.Comp : null;

  const slotTop = place === 'top' && AnnComp ? <AnnComp cfg={cfg} /> : null;

  return (
    <div className={cfg.motion === 'off' ? '' : 'motion-on'}
      style={{
        background: "radial-gradient(23.86% 33.62% at 51% 48%, rgba(255,0,255,0.04) 0%, rgba(172,143,255,0.04) 100%), #111111",
        color: 'var(--sy-text)', minHeight: '100%',
      }}>
      <div style={{
        backgroundImage: "url('assets/landingBg.svg')",
        backgroundSize: '760px', backgroundPosition: 'center 150px', backgroundRepeat: 'no-repeat',
      }}>
        <SynapseNav slotTop={slotTop} />

        {/* band-at-top slot — full-width, directly under the nav, leading the page */}
        {place === 'belowNav' && AnnComp && <AnnComp cfg={cfg} />}

        <main style={{ padding: '0 clamp(16px,4vw,24px)' }}>
          {/* HERO */}
          <section style={{ padding: '40px 0 28px', textAlign: 'center' }}>
            <h1 style={{ margin: '0 auto', maxWidth: 600, fontSize: 'clamp(36px,6vw,48px)', fontWeight: 500, color: '#fff', letterSpacing: '-0.02em', lineHeight: 1.05 }}>
              Secure cross-chain communication
            </h1>
            <p style={{ margin: '18px auto 0', maxWidth: 520, color: 'var(--sy-secondary)', fontSize: 17, lineHeight: 1.55 }}>
              Build truly cross-chain applications using the Synapse Protocol — the most widely used interoperability layer in crypto.
            </p>
            <div style={{ display: 'flex', justifyContent: 'center', gap: 12, marginTop: 28, flexWrap: 'wrap' }}>
              <a href="#" style={{ display: 'inline-flex', alignItems: 'center', height: 48, padding: '0 20px', borderRadius: 8, border: '1px solid #fff', background: '#2f2f2f', color: '#fff', fontWeight: 500, fontSize: 15 }}>Read the docs</a>
              <a href="#" style={{ display: 'inline-flex', alignItems: 'center', height: 48, padding: '0 20px', borderRadius: 10, border: '1px solid var(--sy-purple)', color: '#fff', fontWeight: 500, fontSize: 15, background: 'linear-gradient(310deg, rgba(255,0,255,0.2), rgba(172,143,255,0.2))' }}>Enter Bridge</a>
            </div>

            {/* hero callout slot */}
            {place === 'hero' && AnnComp && (
              <div style={{ marginTop: 40 }}><AnnComp cfg={cfg} /></div>
            )}
          </section>

          {/* afterHero band slot (full-bleed) */}
          {place === 'afterHero' && AnnComp && (
            <div style={{ margin: '8px calc(50% - 50vw) 8px', width: '100vw' }}>
              <AnnComp cfg={cfg} />
            </div>
          )}

          {/* SECURITY */}
          <section style={{ maxWidth: 960, margin: '0 auto', padding: '56px 0' }}>
            <div style={{ textAlign: 'center', maxWidth: 600, margin: '0 auto 44px' }}>
              <SectionLabel>Security</SectionLabel>
              <h2 style={{ margin: 0, fontSize: 'clamp(28px,4vw,38px)', fontWeight: 500, color: '#fff', letterSpacing: '-0.02em' }}>The most secure way to go cross-chain</h2>
            </div>
            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit,minmax(240px,1fr))', gap: 32 }}>
              {SUPPORT_CARDS.map((c, i) => <SupportCard key={c.t} icon={i} {...c} />)}
            </div>
          </section>

          {/* STATS band */}
          <section style={{ margin: '8px calc(50% - 50vw)', width: '100vw', background: 'var(--sy-bg-light)' }}>
            <div style={{ maxWidth: 960, margin: '0 auto', padding: '44px clamp(18px,4vw,40px)' }}>
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', gap: 20, flexWrap: 'wrap', marginBottom: 24 }}>
                <div>
                  <h2 style={{ margin: 0, fontSize: 28, fontWeight: 500, color: '#fff' }}>Battle tested</h2>
                  <p style={{ margin: '6px 0 0', color: 'var(--sy-secondary)' }}>Synapse has processed billions in cross-chain volume.</p>
                </div>
                <a href="#" style={{ display: 'inline-flex', alignItems: 'center', borderRadius: 10, border: '1px solid var(--sy-purple)', padding: '11px 18px', color: '#fff', fontSize: 14, background: 'linear-gradient(310deg, rgba(255,0,255,0.2), rgba(172,143,255,0.2))' }}>Go to Explorer</a>
              </div>
              <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit,minmax(180px,1fr))', gap: 16 }}>
                {[['Total Value Locked', '$142.8M'], ['Total Bridge Volume', '$47.3B'], ['Total TX Count', '12.6M']].map(([t, v]) => (
                  <div key={t} style={{ padding: 16 }}>
                    <div style={{ color: '#fff', opacity: 0.75, fontSize: 15 }}>{t}</div>
                    <div style={{ fontSize: 30, fontWeight: 500, color: '#fff', marginTop: 4 }}>{v}</div>
                  </div>
                ))}
              </div>
            </div>
          </section>

          {/* BRIDGE */}
          <section style={{ maxWidth: 960, margin: '0 auto', padding: '56px 0' }}>
            <div style={{ maxWidth: 520, marginBottom: 36 }}>
              <SectionLabel>The Bridge</SectionLabel>
              <h2 style={{ margin: 0, fontSize: 'clamp(28px,4vw,38px)', fontWeight: 500, color: '#fff', letterSpacing: '-0.02em' }}>Powering the cross-chain economy</h2>
              <p style={{ margin: '14px 0 0', color: 'var(--sy-secondary)', fontSize: 16, lineHeight: 1.6 }}>
                <strong style={{ color: '#fff', fontWeight: 500 }}>The Synapse Bridge</strong> is built on top of the most widely used cross-chain messaging protocol.
              </p>
            </div>
            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit,minmax(240px,1fr))', gap: 32 }}>
              {BRIDGE_CARDS.map((c, i) => <SupportCard key={c.t} icon={i} {...c} />)}
            </div>
          </section>

          {/* INTEGRATION chains */}
          <section style={{ maxWidth: 1100, margin: '0 auto', padding: '40px 0 72px' }}>
            <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', gap: 24, flexWrap: 'wrap', marginBottom: 36 }}>
              <h2 style={{ margin: 0, paddingRight: 24, borderRight: '1px solid rgba(255,255,255,0.25)', fontSize: 26, color: '#fff', fontWeight: 500 }}>Widely integrated</h2>
              <p style={{ margin: 0, maxWidth: 460, color: 'var(--sy-secondary)' }}>Synapse is supported across 20+ chains and the apps building on top of them.</p>
            </div>
            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit,minmax(150px,1fr))', gap: 12, maxWidth: 920, margin: '0 auto' }}>
              {CHAINS.map(([name, icon]) => (
                <a key={name} href="#" style={{ border: '1px solid var(--sy-panel)', padding: '18px 8px', textAlign: 'center', transition: 'border-color .15s' }}
                  onMouseEnter={(e) => (e.currentTarget.style.borderColor = 'rgba(255,255,255,0.5)')}
                  onMouseLeave={(e) => (e.currentTarget.style.borderColor = 'var(--sy-panel)')}>
                  <img src={`assets/chains/${icon}.svg`} alt="" width="44" height="44" style={{ borderRadius: '50%', marginBottom: 8 }} />
                  <div style={{ fontSize: 16, color: '#fff', fontWeight: 500 }}>{name}</div>
                  <div style={{ fontSize: 13, color: 'var(--sy-secondary)', opacity: 0.75, marginTop: 2 }}>Layer {name === 'Ethereum' || name === 'Avalanche' || name === 'BNB Chain' ? 1 : 2}</div>
                </a>
              ))}
            </div>
          </section>
        </main>

        {/* FOOTER */}
        <footer style={{ borderTop: '1px solid rgba(255,255,255,0.08)', padding: '40px clamp(18px,4vw,40px) 104px' }}>
          <div style={{ maxWidth: 1100, margin: '0 auto', display: 'flex', justifyContent: 'space-between', alignItems: 'center', gap: 20, flexWrap: 'wrap' }}>
            <img src="assets/synapse-logo.svg" alt="Synapse by Hypercall" style={{ height: 32 }} />
            <div style={{ display: 'flex', gap: 22, color: 'var(--sy-secondary)', fontSize: 14, flexWrap: 'wrap' }}>
              <a href="#" style={{ color: 'var(--sy-secondary)' }}>Terms of Use</a>
              <a href="#" style={{ color: 'var(--sy-secondary)' }}>Privacy Policy</a>
              <span>© 2026 Synapse</span>
            </div>
          </div>
        </footer>

        {/* fixed floating slot */}
        {place === 'fixed' && AnnComp && <AnnComp cfg={cfg} />}
      </div>
    </div>
  );
}

Object.assign(window, { SynapsePage, SynapseNav });
