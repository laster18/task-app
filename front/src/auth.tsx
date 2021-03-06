import { NextPageContext } from 'next'
import { parseCookies } from 'nookies'
import Router from 'next/router'
import cookie from 'js-cookie'

const AFTER_LOGIN_URL = '/'

export interface Auth {
  jwt?: string
}

export function loadAuthFromCookie(ctx: NextPageContext): Auth {
  // const { token, jwt } = parseCookies(ctx)
  const { jwt } = parseCookies(ctx)
  return { jwt }
}

export const login = ({ token }: { token: string }) => {
  cookie.set('token', token, { expires: 1 })
  Router.push(AFTER_LOGIN_URL)
}
