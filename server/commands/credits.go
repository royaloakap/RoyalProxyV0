package commands

import (
	"fmt"
	"webserver/src/server/sessions"
	"strings"
)

func Credits(session *sessions.Session, args []string) {
	// Informations sur l'auteur et la source
	author := ""
	sourceLines := "\u001B[38;5;041 m< 2700 \u001B[38;5;230m"

	// Create a builder for constructing the message
	var builder strings.Builder

	// Append basic credits information
	builder.WriteString("[+]\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*[+]\r\n")
	builder.WriteString("\u001B[38;5;123mCredits:\r\n")
	builder.WriteString(fmt.Sprintf("\u001B[38;5;230mAuthor: \u001B[38;5;44m t.me/Royaloakap / t.me/Royal_FAQ%s\r\n", author))
	builder.WriteString(fmt.Sprintf("\u001B[38;5;220mVersion: \u001B[38;5;207m 1.1 \u001B[38;5;230m %s\r\n", sourceLines))
	builder.WriteString("discord.gg/RoyalC2.\r\n")

	// Append additional information
	builder.WriteString("\u001B[38;5;255mRoyal CNC Free Version is a custom written source\r\n")
	builder.WriteString("with less than 2,700 lines of code.\r\n")
	builder.WriteString("\u001B[38;5;255mThis Src was developed solely by https://t.me/Royaloakap\r\n")
	builder.WriteString("\u001B[38;5;255mThank you for using My CNC ! \u001B[38;5;196m♥ \u001B[38;5;255mt.me/RoyalSRC\r\n")
	builder.WriteString("[+]\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*\u001B[38;5;196m♥\u001B[38;5;230m*[+]\r\n")

	// Print the message
	fmt.Fprintf(session.Conn, builder.String())
}
