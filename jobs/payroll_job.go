package jobs

import (
	"fmt"
	"strconv"
	"time"

	apset "github.com/ebikode/eLearning-core/domain/app_setting"
	pyr "github.com/ebikode/eLearning-core/domain/payroll"
	tx "github.com/ebikode/eLearning-core/domain/tax"
	emp "github.com/ebikode/eLearning-core/domain/user"
	md "github.com/ebikode/eLearning-core/model"
	ut "github.com/ebikode/eLearning-core/utils"
)

var users []*md.User

// RunDefaultPayrollGenerationJob ...
func RunDefaultPayrollGenerationJob(
	pys pyr.PayrollService,
	emps emp.UserService, txs tx.TaxService,
) {
	count := 5

	for i := 0; i < count; i++ {
		month := uint(i) + 1
		year := uint(2020)
		payroll := pys.GetSinglePayrollByMonthYear(month, year)

		if payroll == nil {
			users := emps.GetAllActivePubUser()

			fmt.Println("Default Payroll Generation Started")

			for _, v := range users {

				if v.Salary != nil {
					netSalary, tax := generateNetSalaryAndTaxDeductions(v.Salary)

					payroll := md.Payroll{
						UserID:        v.ID,
						GrossSalary:   v.Salary.Salary,
						NetSalary:     netSalary,
						Month:         month,
						Year:          year,
						PaymentStatus: ut.Success,
						Status:        ut.Approved,
					}

					newPayroll, _, err := pys.CreatePayroll(payroll)

					if err == nil {
						tax.PayrollID = newPayroll.ID
						txs.CreateTax(tax)
					}

				}

			}

			fmt.Println("Default Payroll Generation Ended")
		}
	}

}

// RunPayrollGenerationJob - Automated Payroll Generation
func RunPayrollGenerationJob(
	pys pyr.PayrollService, aps apset.AppSettingService,
	emps emp.UserService, txs tx.TaxService,
) {

	genDay := aps.GetAppSettingByKey(ut.PayrollGenerationDayKey, "admin")

	generationDate, _ := strconv.ParseInt(genDay.Value, 10, 64)

	now := time.Now()

	todayDate := now.Day()
	todayMonth := uint(int(now.Month()))
	todayYear := uint(now.Year())

	if todayDate >= int(generationDate) {
		checkPayroll := pys.GetSinglePayrollByMonthYear(todayMonth, todayYear)

		if checkPayroll == nil {
			users := emps.GetAllActivePubUser()

			fmt.Println("Payroll Generation Automation Started")

			for _, v := range users {

				if v.Salary != nil {

					netSalary, tax := generateNetSalaryAndTaxDeductions(v.Salary)

					payroll := md.Payroll{
						UserID:      v.ID,
						GrossSalary: v.Salary.Salary,
						NetSalary:   netSalary,
						Month:       todayMonth,
						Year:        todayYear,
					}

					newPayroll, _, err := pys.CreatePayroll(payroll)

					if err == nil {
						tax.PayrollID = newPayroll.ID
						txs.CreateTax(tax)
					}

				}

			}

			fmt.Println("Payroll Generation Automation Ended")

		}
	}

}

// RunPayrollPaymentJob - Automated Payroll Payment
func RunPayrollPaymentJob(pys pyr.PayrollService, aps apset.AppSettingService) {

	payDay := aps.GetAppSettingByKey(ut.PayDayKey, "admin")

	payDate, _ := strconv.ParseInt(payDay.Value, 10, 64)

	now := time.Now()

	todayDate := now.Day()
	todayMonth := uint(int(now.Month()))
	todayYear := uint(now.Year())

	if todayDate >= int(payDate) {
		checkPayroll := pys.GetSinglePayrollByMonthYear(todayMonth, todayYear)

		if checkPayroll != nil && checkPayroll.PaymentStatus == ut.Pending && checkPayroll.Status == ut.Approved {

			fmt.Println("Payroll Payment Automation Started")

			pys.UpdatePayrollPaymentStatus(ut.Success, int(todayMonth), int(todayYear))

			fmt.Println("Payroll Payment Automation Ended")

		}
	}

}

func generateNetSalaryAndTaxDeductions(salary *md.Salary) (float64, md.Tax) {

	GrossSalary := salary.Salary

	tax := md.Tax{
		Pension: percentage(salary.Pension) * GrossSalary,
		Paye:    percentage(salary.Paye) * GrossSalary,
		Nsitf:   percentage(salary.Nsitf) * GrossSalary,
		Nhf:     percentage(salary.Nhf) * GrossSalary,
		Itf:     percentage(salary.Itf) * GrossSalary,
	}

	deductions := tax.Pension + tax.Paye + tax.Nsitf + tax.Nhf + tax.Itf

	netSalary := GrossSalary - deductions

	return netSalary, tax

}

func percentage(value float64) float64 {
	return (value / 100)
}
