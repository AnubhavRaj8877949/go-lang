package main

import "fmt"

// ==========================================================
// 1. COMPONENT INTERFACE (Base Rule)
// ==========================================================
type Notifier interface {
	Send(message string)
}

// ==========================================================
// 2. CONCRETE COMPONENT (Asli Base Kaam - Email)
// ==========================================================
type EmailNotifier struct {
	emailAddress string
}

func (e *EmailNotifier) Send(message string) {
	fmt.Printf("📧 [EMAIL to %s]: %s\n", e.emailAddress, message)
}

// ==========================================================
// 3. DECORATOR 1: SMS Wrapper
// ==========================================================
type SMSDecorator struct {
	wrapped     Notifier // Apne hi interface ko andar wrap kiya
	phoneNumber string
}

func (s *SMSDecorator) Send(message string) {
	// Step 1: Pehle andar wale notifier ka kaam chalaya
	s.wrapped.Send(message)
	// Step 2: Phir apna extra feature (SMS) add kiya
	fmt.Printf("📱 [SMS to %s]: %s\n", s.phoneNumber, message)
}

// ==========================================================
// 4. DECORATOR 2: Slack Wrapper
// ==========================================================
type SlackDecorator struct {
	wrapped      Notifier // Apne hi interface ko andar wrap kiya
	slackChannel string
}

func (sl *SlackDecorator) Send(message string) {
	// Step 1: Pehle andar wale notifier ka kaam chalaya
	sl.wrapped.Send(message)
	// Step 2: Phir apna extra feature (Slack) add kiya
	fmt.Printf("💬 [SLACK #%s]: %s\n", sl.slackChannel, message)
}

// ==========================================================
// 5. CLIENT (Ise sirf 'Notifier' interface se matlab hai)
// ==========================================================
type AlertManager struct{}

func (am *AlertManager) TriggerAlert(n Notifier, msg string) {
	fmt.Println("--- Alert Triggered ---")
	n.Send(msg)
	fmt.Println()
}

// ==========================================================
// 6. MAIN EXECUTION (Saare Cases Ek Saath)
// ==========================================================
func main() {
	manager := &AlertManager{}

	// Base Object: Sirf Email
	baseEmail := &EmailNotifier{emailAddress: "admin@company.com"}

	// CASE 1: Normal User (Sirf Email chahiye)
	fmt.Println("CASE 1: Basic Notification")
	manager.TriggerAlert(baseEmail, "Daily backup completed.")

	// CASE 2: Important Alert (Email + SMS)
	fmt.Println("CASE 2: Email + SMS Notification")
	emailWithSMS := &SMSDecorator{
		wrapped:     baseEmail,
		phoneNumber: "+91-9876543210",
	}
	manager.TriggerAlert(emailWithSMS, "High CPU usage detected!")

	// CASE 3: Team Alert (Email + Slack)
	fmt.Println("CASE 3: Email + Slack Notification")
	emailWithSlack := &SlackDecorator{
		wrapped:      baseEmail,
		slackChannel: "devops-alerts",
	}
	manager.TriggerAlert(emailWithSlack, "New deployment started.")

	// CASE 4: Critical Emergency (Email + SMS + Slack teeno ek saath!)
	fmt.Println("CASE 4: Full Stack (Email + SMS + Slack)")
	fullAlert := &SlackDecorator{
		wrapped: &SMSDecorator{
			wrapped:     baseEmail,
			phoneNumber: "+91-9876543210",
		},
		slackChannel: "critical-incidents",
	}
	manager.TriggerAlert(fullAlert, "CRITICAL: Production Database Down!")
}
