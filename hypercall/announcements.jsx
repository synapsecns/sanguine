/* announcements.jsx
 * The Hypercall announcement, in four placements × tunable loudness/identity/motion.
 * Shared by both the live page rebuild and the side-by-side canvas.
 * Exports to window: HC_LINK, HypercallWordmark, accentFor, ANNOUNCEMENTS
 */

const HC_LINK = 'https://app.hypercall.xyz/';

/* "by Synapse" — plain, unstyled */
function SynByline({ size = 13 }) {
  return (
    <span style={{ fontSize: size, color: 'var(--sy-secondary)', whiteSpace: 'nowrap', fontWeight: 400 }}>
      by Synapse
    </span>
  );
}

/* Hypercall wordmark — paths lifted from the repo (Hypercall.tsx), recolored to currentColor */
function HypercallWordmark({ height = 18, style }) {
  return (
    <svg viewBox="0 0 180 36" height={height} width={(180 / 36) * height}
      fill="none" xmlns="http://www.w3.org/2000/svg"
      style={{ display: 'block', ...style }} aria-label="Hypercall">
      <g fill="currentColor">
        <path d="M169.741 28.5L174.221 0.179993H179.261L174.781 28.5H169.741Z" />
        <path d="M174 24.5H177L176.5 28.5H173.5L174 24.5Z" />
        <path d="M172.5 0.174988H175.5L175 4.17499H172L172.5 0.174988Z" />
        <path d="M158.819 28.5L163.299 0.179993H168.339L163.859 28.5H158.819Z" />
        <path d="M163 24.5H166L165.5 28.5H162.5L163 24.5Z" />
        <path d="M161.5 0.174988H164.5L164 4.17499H161L161.5 0.174988Z" />
        <path d="M133.25 19.46C133.25 16 134.75 12.48 137.25 10.25C139.75 8.02002 142.661 8.02002 145.341 8.02002C150.421 8.02002 153.261 11.46 153.261 16.58C153.261 23.66 148.581 28.82 142.021 28.82C139.461 28.82 137.25 27.25 136 26.25C134.75 25.25 133.25 22.92 133.25 19.46ZM143.601 24.46C147.121 24.46 149.681 21.54 149.681 17.38C149.681 14.38 148.001 12.38 145.201 12.38C141.521 12.38 138.801 15.3 138.801 19.46C138.801 22.46 140.641 24.46 143.601 24.46Z" />
        <path d="M156.341 8.33997H151.301L148.097 28.5H153.137L156.341 8.33997Z" />
        <path d="M152.5 24.5H155.5L155 28.5H152L152.5 24.5Z" />
        <path d="M120.281 28.82C114.881 28.82 110.961 24.9 110.961 19.62C110.961 12.86 115.801 8.02002 122.641 8.02002C127.681 8.02002 131.321 11.06 131.401 15.3L126.241 16.18C126.201 14.02 124.481 12.38 122.121 12.38C118.561 12.38 116.281 15.1 116.281 19.18C116.281 22.22 118.081 24.46 120.761 24.46C123.161 24.46 124.921 23.18 125.441 21.06L130.361 22.06C129.361 26.18 125.521 28.82 120.281 28.82Z" />
        <path d="M101.215 21.1001L100.055 28.5001H95.0146L98.2147 8.25H103.255H111L110.25 12.1401H102.655L101.215 21.1001Z" />
        <path d="M93.0411 16.42C93.0411 17.46 92.8411 18.7 92.5611 19.66H77.6011C77.7211 22.78 79.4811 24.58 82.4411 24.58C84.7211 24.58 86.4011 23.5 87.0011 21.5L91.7611 22.78C90.6011 26.38 86.6411 28.82 81.8011 28.82C76.4011 28.82 72.6411 25.06 72.6411 19.62C72.6411 13.98 76.4411 8.02002 83.9611 8.02002C90.0011 8.02002 93.0411 12.22 93.0411 16.42ZM87.8411 16.06C88.0011 13.66 86.2811 11.86 83.3611 11.86C80.7211 11.86 78.8811 13.54 78.0811 16.06H87.8411Z" />
        <path d="M46.8625 35.78L49.4557 19.46C49.4557 19.46 50 13.5 53 10.75C56 8.00003 59.5 7.75001 62.25 8.00001C66 8.25001 70.1826 11.46 70.1826 16.58C70.1826 23.66 65.5026 28.82 58.9426 28.82C56.3826 28.82 54.3426 27.9 53.3826 26.26L51.9026 35.78H46.8625ZM59.0226 24.46C62.5426 24.46 65.1026 21.54 65.1026 17.38C65.1026 14.38 63.4226 12.38 60.6226 12.38C56.9426 12.38 54.2226 15.3 54.2226 19.46C54.2226 22.46 56.0625 24.46 59.0226 24.46Z" />
        <path d="M45.3174 8.34003H50.3174L35.3574 35.78H30.5174L35.0374 27.42L30.3574 8.34003H35.5574L38.3174 21.3L45.3174 8.34003Z" />
        <path d="M34.5 31.78H40.5L40 35.78H34L34.5 31.78Z" />
        <path d="M29.5 8.34998H35.5L35 12.35H29L29.5 8.34998Z" />
        <path d="M23.2802 0.5H28.7202L24.2802 28.5H18.8402L20.7602 16.3H8.52016L6.60016 28.5H1.16016L5.60016 0.5H11.0402L9.28016 11.62H21.5202L23.2802 0.5Z" />
        <path d="M3 0.5H6L5.5 4.5H2.5L3 0.5Z" />
        <path d="M24.5 24.5H27.5L27 28.5H24L24.5 24.5Z" />
      </g>
    </svg>
  );
}

/* identity → palette. cobrand = Synapse gradient surface + Hypercall lime CTA */
function accentFor(identity) {
  const lime = 'var(--hc-lime)';
  const grad = 'linear-gradient(95deg, var(--sy-pink), var(--sy-purple))';
  if (identity === 'synapse') {
    return {
      mark: 'gradient',          // render synapse-tinted wordmark
      markColor: '#ffffff',
      rule: grad,
      glow: 'rgba(172,143,255,0.18)',
      ctaBg: grad,
      ctaColor: '#ffffff',
      ctaBorder: 'transparent',
      chip: 'rgba(172,143,255,0.14)',
      chipBorder: 'rgba(172,143,255,0.32)',
      chipText: '#d9ccff',
      dot: 'var(--sy-purple)',
      accentText: 'var(--sy-purple)',
    };
  }
  if (identity === 'hypercall') {
    return {
      mark: 'lime',
      markColor: lime,
      rule: 'linear-gradient(95deg, var(--hc-lime), var(--hc-green-text))',
      glow: 'rgba(169,250,56,0.16)',
      ctaBg: lime,
      ctaColor: 'var(--hc-ink)',
      ctaBorder: 'transparent',
      chip: 'rgba(169,250,56,0.12)',
      chipBorder: 'rgba(169,250,56,0.4)',
      chipText: 'var(--hc-green-text)',
      dot: lime,
      accentText: lime,
    };
  }
  /* cobrand (default) */
  return {
    mark: 'lime',
    markColor: lime,
    rule: grad,
    glow: 'rgba(169,250,56,0.14)',
    ctaBg: lime,
    ctaColor: 'var(--hc-ink)',
    ctaBorder: 'transparent',
    chip: 'rgba(169,250,56,0.12)',
    chipBorder: 'rgba(169,250,56,0.4)',
    chipText: 'var(--hc-green-text)',
    dot: lime,
    accentText: lime,
  };
}

/* loudness → sizing knobs */
function loudFor(loud) {
  if (loud === 'subtle') return { padY: 10, font: 14, h: 'auto', headFont: 15, subFont: 13, ctaPad: '8px 15px', ctaFont: 13.5, surfaceAlpha: 0.5 };
  if (loud === 'bold')   return { padY: 18, font: 16.5, h: 'auto', headFont: 21, subFont: 15.5, ctaPad: '13px 24px', ctaFont: 16, surfaceAlpha: 1 };
  return { padY: 14, font: 15.5, h: 'auto', headFont: 18, subFont: 14.5, ctaPad: '11px 20px', ctaFont: 15, surfaceAlpha: 0.78 };
}

function CTA({ cfg, a, l, children }) {
  return (
    <a href={HC_LINK} target="_blank" rel="noreferrer"
      className={cfg.motion === 'lively' || cfg.motion === 'subtle' ? 'hc-cta-sheen' : ''}
      style={{
        display: 'inline-flex', alignItems: 'center', gap: 8,
        background: a.ctaBg, color: a.ctaColor,
        border: `1px solid ${a.ctaBorder}`,
        padding: l.ctaPad, borderRadius: 10,
        fontWeight: 600, fontSize: l.ctaFont, letterSpacing: '-0.01em',
        whiteSpace: 'nowrap', cursor: 'pointer',
        boxShadow: cfg.identity !== 'synapse' ? '0 6px 22px -10px rgba(169,250,56,0.6)' : '0 6px 22px -10px rgba(172,143,255,0.6)',
      }}>
      {children || cfg.ctaLabel}
      <span aria-hidden="true">→</span>
    </a>
  );
}

function ProofChips({ cfg, a }) {
  if (!cfg.proof) return null;
  return (
    <div style={{ display: 'flex', alignItems: 'center', gap: 8, flexWrap: 'wrap' }}>
      <span style={{
        display: 'inline-flex', alignItems: 'center', gap: 7,
        fontFamily: 'var(--mono)', fontSize: 11.5, letterSpacing: '0.02em', whiteSpace: 'nowrap',
        color: a.chipText, background: a.chip,
        border: `1px solid ${a.chipBorder}`, borderRadius: 999, padding: '4px 11px',
      }}>
        <img src="assets/hyperliquid.svg" alt="" width="13" height="13" style={{ borderRadius: '50%' }} />
        Built on Hyperliquid
      </span>
      <span style={{
        display: 'inline-flex', alignItems: 'center', gap: 7,
        fontFamily: 'var(--mono)', fontSize: 11.5, letterSpacing: '0.02em', whiteSpace: 'nowrap',
        color: a.chipText, background: a.chip,
        border: `1px solid ${a.chipBorder}`, borderRadius: 999, padding: '4px 11px',
      }}>
        <img src="assets/spcx.svg" alt="" width="13" height="13" style={{ borderRadius: '50%' }} />
        SpaceX live now
      </span>
    </div>
  );
}

function DismissX({ cfg, color }) {
  if (!cfg.dismissible) return null;
  return (
    <button onClick={cfg.onDismiss} aria-label="Dismiss"
      style={{
        flexShrink: 0, width: 26, height: 26, display: 'grid', placeItems: 'center',
        background: 'transparent', border: 'none', cursor: 'pointer',
        color: color || 'rgba(255,255,255,0.45)', borderRadius: 6, padding: 0,
      }}
      onMouseEnter={(e) => (e.currentTarget.style.color = '#fff')}
      onMouseLeave={(e) => (e.currentTarget.style.color = color || 'rgba(255,255,255,0.45)')}>
      <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round">
        <path d="M6 6l12 12M18 6 6 18" />
      </svg>
    </button>
  );
}

/* ============================================================
 * 1 · BANNER — full-width strip across the very top of the page
 * ============================================================ */
function AnnBanner({ cfg }) {
  const a = accentFor(cfg.identity);
  const l = loudFor(cfg.loud);
  return (
    <div style={{ position: 'relative', overflow: 'hidden' }}>
      {/* gradient hairline rule under the bar */}
      <div style={{ position: 'absolute', left: 0, right: 0, bottom: 0, height: 1.5, background: a.rule, opacity: 0.9 }} />
      <div style={{
        background: cfg.loud === 'bold'
          ? `linear-gradient(90deg, rgba(255,0,153,0.16), rgba(172,143,255,0.16) 55%, rgba(169,250,56,0.10))`
          : `linear-gradient(90deg, rgba(255,0,153,0.10), rgba(172,143,255,0.10))`,
        backgroundSize: cfg.motion === 'lively' ? '200% 100%' : 'auto',
        animation: cfg.motion === 'lively' ? 'hc-gradient-pan 9s linear infinite alternate' : 'none',
      }}>
        <div style={{
          maxWidth: 1180, margin: '0 auto',
          padding: `${l.padY}px clamp(16px,4vw,40px)`,
          display: 'flex', alignItems: 'center', gap: 'clamp(12px,2vw,22px)', flexWrap: 'wrap',
        }}>
          <span className="hc-live-dot" style={{ background: a.dot, flexShrink: 0 }} />
          <div style={{ display: 'flex', alignItems: 'baseline', gap: 9, flexWrap: 'wrap', flex: 1, minWidth: 220 }}>
            <span style={{ display: 'inline-flex', alignItems: 'baseline', gap: 8 }}>
              <span style={{ display: 'inline-flex', color: a.markColor, transform: 'translateY(3px)' }}>
                <HypercallWordmark height={cfg.loud === 'bold' ? 19 : 16} />
              </span>
              <SynByline size={cfg.loud === 'bold' ? 14 : 13} />
            </span>
            <span style={{ fontSize: l.font, color: 'var(--sy-text)', fontWeight: 500, letterSpacing: '-0.01em' }}>
              is live — options on anything,
              <span style={{ color: a.accentText }}> starting with SpaceX.</span>
            </span>
            <span style={{ fontSize: l.font - 1.5, color: 'var(--sy-secondary)' }}>
              Defined risk. No liquidations. On Hyperliquid.
            </span>
          </div>
          <CTA cfg={cfg} a={a} l={l} />
          <DismissX cfg={cfg} />
        </div>
      </div>
    </div>
  );
}

/* ============================================================
 * 2 · HERO CALLOUT — a framed card sitting inside the hero column
 * ============================================================ */
function AnnHeroCallout({ cfg }) {
  const a = accentFor(cfg.identity);
  const l = loudFor(cfg.loud);
  return (
    <div style={{
      position: 'relative', maxWidth: 620, margin: '0 auto',
      borderRadius: 18, overflow: 'hidden',
      animation: cfg.motion !== 'off' ? 'hc-rise 0.6s cubic-bezier(.2,.7,.3,1)' : 'none',
    }}>
      <div style={{ position: 'absolute', inset: 0, padding: 1.5, borderRadius: 18,
        background: a.rule, WebkitMask: 'linear-gradient(#000 0 0) content-box, linear-gradient(#000 0 0)',
        WebkitMaskComposite: 'xor', maskComposite: 'exclude', pointerEvents: 'none' }} />
      <div style={{
        background: `radial-gradient(120% 140% at 0% 0%, ${a.glow}, transparent 60%), rgba(25,22,28,0.92)`,
        backdropFilter: 'blur(8px)',
        padding: cfg.loud === 'bold' ? '24px 26px' : '20px 22px',
        display: 'flex', alignItems: 'center', gap: 20, flexWrap: 'wrap',
      }}>
        <div style={{ flex: 1, minWidth: 240 }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 12 }}>
            <span className="hc-live-dot" style={{ background: a.dot }} />
            <span style={{ fontFamily: 'var(--mono)', fontSize: 11, letterSpacing: '0.14em', textTransform: 'uppercase', color: a.accentText }}>
              New · A Synapse product
            </span>
          </div>
          <div style={{ display: 'flex', alignItems: 'baseline', gap: 8, flexWrap: 'wrap' }}>
            <span style={{ color: a.markColor, display: 'inline-flex', transform: 'translateY(3px)' }}>
              <HypercallWordmark height={cfg.loud === 'bold' ? 22 : 19} />
            </span>
            <span style={{ fontSize: l.headFont + 2, fontWeight: 600, color: '#fff', letterSpacing: '-0.02em' }}>
              is live.
            </span>
            <SynByline size={14} />
          </div>
          <p style={{ margin: '8px 0 16px', fontSize: l.subFont, lineHeight: 1.5, color: 'var(--sy-secondary)', maxWidth: 360 }}>
            Options on anything — a call on SpaceX, a put on BTC. Defined risk, no liquidations, settled on Hyperliquid.
          </p>
          <div style={{ display: 'flex', alignItems: 'center', gap: 14, flexWrap: 'wrap' }}>
            <CTA cfg={cfg} a={a} l={l} />
            <ProofChips cfg={cfg} a={a} />
          </div>
        </div>
        <DismissX cfg={cfg} />
      </div>
    </div>
  );
}

/* ============================================================
 * 3 · DEDICATED BAND — full-bleed section between hero & content
 * ============================================================ */
function AnnBand({ cfg }) {
  const a = accentFor(cfg.identity);
  const l = loudFor(cfg.loud);
  return (
    <section style={{ position: 'relative', overflow: 'hidden',
      borderTop: '1px solid rgba(255,255,255,0.06)', borderBottom: '1px solid rgba(255,255,255,0.06)',
      background: 'linear-gradient(180deg, #141218, #0f0d12)' }}>
      {/* glow */}
      <div style={{ position: 'absolute', top: '-40%', left: '50%', transform: 'translateX(-50%)',
        width: 900, height: 460, background: `radial-gradient(45% 50% at 50% 40%, ${a.glow}, transparent 70%)`, pointerEvents: 'none' }} />
      {/* moving gradient top rule */}
      <div style={{ position: 'absolute', top: 0, left: 0, right: 0, height: 2, background: a.rule,
        backgroundSize: '200% 100%', animation: cfg.motion === 'lively' ? 'hc-gradient-pan 7s linear infinite alternate' : 'none' }} />
      <div style={{ position: 'relative', maxWidth: 1100, margin: '0 auto',
        padding: cfg.loud === 'bold' ? '56px clamp(18px,4vw,40px)' : '44px clamp(18px,4vw,40px)',
        display: 'flex', alignItems: 'center', justifyContent: 'space-between', gap: 36, flexWrap: 'wrap' }}>
        <div style={{ flex: 1, minWidth: 300 }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 16 }}>
            <span className="hc-live-dot" style={{ background: a.dot }} />
            <span style={{ fontFamily: 'var(--mono)', fontSize: 12, letterSpacing: '0.14em', textTransform: 'uppercase', color: a.accentText }}>
              Now live on mainnet
            </span>
          </div>
          <h2 style={{ margin: 0, fontSize: 'clamp(28px,4.4vw,42px)', lineHeight: 1.02, letterSpacing: '-0.03em', color: '#fff', fontWeight: 500 }}>
            Synapse has built{' '}
            <span style={{ color: a.markColor, display: 'inline-flex', verticalAlign: 'baseline', transform: 'translateY(4px)' }}>
              <HypercallWordmark height={cfg.loud === 'bold' ? 38 : 32} />
            </span>
          </h2>
          <p style={{ margin: '16px 0 0', fontSize: l.subFont + 2, lineHeight: 1.55, color: 'var(--sy-secondary)', maxWidth: 540 }}>
            Options on anything — fractional, defined-risk options, settled on Hyperliquid. Same team, same token; now trading SpaceX, with crypto and equities next. No liquidations, open nights &amp; weekends.
          </p>
          <div style={{ marginTop: 22 }}><ProofChips cfg={cfg} a={a} /></div>
        </div>
        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start', gap: 14 }}>
          <CTA cfg={cfg} a={a} l={l}>{cfg.ctaLabel}</CTA>
          <a href={HC_LINK} target="_blank" rel="noreferrer"
            style={{ fontSize: 13.5, color: 'var(--sy-secondary)', borderBottom: '1px solid rgba(255,255,255,0.18)', paddingBottom: 1 }}>
            See how it works
          </a>
        </div>
        <div style={{ position: 'absolute', top: 14, right: 14 }}><DismissX cfg={cfg} /></div>
      </div>
    </section>
  );
}

/* ============================================================
 * 4 · FLOATING CARD — sticky bottom-right, appears on scroll
 * ============================================================ */
function AnnFloating({ cfg }) {
  const a = accentFor(cfg.identity);
  const l = loudFor(cfg.loud);
  return (
    <div style={{
      position: cfg.fixed === false ? 'relative' : 'fixed',
      right: cfg.fixed === false ? 'auto' : 22, bottom: cfg.fixed === false ? 'auto' : 22,
      zIndex: 60, width: cfg.fixed === false ? '100%' : 'min(360px, calc(100vw - 32px))',
      maxWidth: 360,
      borderRadius: 16, overflow: 'hidden',
      boxShadow: '0 30px 60px -24px rgba(0,0,0,0.8)',
      animation: cfg.motion !== 'off' ? 'hc-rise 0.5s cubic-bezier(.2,.7,.3,1)' : 'none',
    }}>
      <div style={{ position: 'absolute', inset: 0, padding: 1.5, borderRadius: 16,
        background: a.rule, WebkitMask: 'linear-gradient(#000 0 0) content-box, linear-gradient(#000 0 0)',
        WebkitMaskComposite: 'xor', maskComposite: 'exclude', pointerEvents: 'none' }} />
      <div style={{ background: `radial-gradient(110% 90% at 100% 0%, ${a.glow}, transparent 55%), rgba(22,19,25,0.97)`,
        backdropFilter: 'blur(10px)', padding: '18px 18px 16px' }}>
        <div style={{ display: 'flex', alignItems: 'flex-start', justifyContent: 'space-between', gap: 10 }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
            <span className="hc-live-dot" style={{ background: a.dot }} />
            <span style={{ fontFamily: 'var(--mono)', fontSize: 10.5, letterSpacing: '0.12em', textTransform: 'uppercase', color: a.accentText }}>
              Now live
            </span>
          </div>
          <DismissX cfg={cfg} />
        </div>
        <div style={{ display: 'flex', alignItems: 'baseline', gap: 8, margin: '12px 0 0', flexWrap: 'wrap' }}>
          <span style={{ color: a.markColor, display: 'inline-flex', transform: 'translateY(3px)' }}>
            <HypercallWordmark height={18} />
          </span>
          <span style={{ fontSize: 16, fontWeight: 600, color: '#fff' }}>· options on anything</span>
          <SynByline size={12.5} />
        </div>
        <p style={{ margin: '9px 0 14px', fontSize: 13.5, lineHeight: 1.5, color: 'var(--sy-secondary)' }}>
          Trade SpaceX, BTC &amp; more. Defined risk, no liquidations, on Hyperliquid.
        </p>
        <div style={{ display: 'flex', alignItems: 'center', gap: 10, flexWrap: 'wrap' }}>
          <CTA cfg={cfg} a={a} l={{ ...l, ctaPad: '11px 18px' }} />
        </div>
      </div>
    </div>
  );
}

const ANNOUNCEMENTS = {
  banner:      { key: 'banner',      label: 'Top banner',      placement: 'top',      Comp: AnnBanner },
  heroCallout: { key: 'heroCallout', label: 'Hero callout',    placement: 'hero',     Comp: AnnHeroCallout },
  band:        { key: 'band',        label: 'Dedicated band',  placement: 'afterHero', Comp: AnnBand },
  bandTop:     { key: 'bandTop',     label: 'Band at top',     placement: 'belowNav', Comp: AnnBand },
  floating:    { key: 'floating',    label: 'Floating card',   placement: 'fixed',    Comp: AnnFloating },
};

Object.assign(window, { HC_LINK, HypercallWordmark, SynByline, accentFor, ANNOUNCEMENTS });
